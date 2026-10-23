package tui

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/mmcdole/gofeed"
)

func testFeed() *gofeed.Feed {
	return &gofeed.Feed{Items: []*gofeed.Item{
		{Title: "First episode", Published: "Mon, 01 Jan 2024", Description: "<p>Hello <b>world</b></p>"},
		{Title: "Second episode", Published: "Tue, 02 Jan 2024", Description: "<p>Another one</p>"},
	}}
}

func newTestModel(t *testing.T) model {
	t.Helper()
	m, err := NewModel(testFeed())
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	return *m
}

func press(t *testing.T, m model, msg tea.Msg) (model, tea.Cmd) {
	t.Helper()
	next, cmd := m.Update(msg)
	nm, ok := next.(model)
	if !ok {
		t.Fatalf("Update returned %T, want model", next)
	}
	return nm, cmd
}

func key(r rune) tea.KeyPressMsg { return tea.KeyPressMsg{Code: r, Text: string(r)} }

func TestNewModelConvertsItems(t *testing.T) {
	m := newTestModel(t)
	items := m.list.Items()
	if len(items) != 2 {
		t.Fatalf("got %d items, want 2", len(items))
	}
	first := items[0].(item)
	if first.title != "First episode" {
		t.Errorf("title = %q", first.title)
	}
	if !strings.Contains(first.desc, "Published at Mon, 01 Jan 2024") {
		t.Errorf("desc missing published line: %q", first.desc)
	}
	if !strings.Contains(first.desc, "**world**") {
		t.Errorf("desc not converted to markdown: %q", first.desc)
	}
}

func TestEnterOpensDetail(t *testing.T) {
	m, _ := press(t, newTestModel(t), tea.KeyPressMsg{Code: tea.KeyEnter})
	if !m.detail || m.choice != "First episode" {
		t.Errorf("detail=%v choice=%q", m.detail, m.choice)
	}
	if !strings.Contains(m.content, "Hello") {
		t.Errorf("content = %q", m.content)
	}
}

func TestBackClearsDetail(t *testing.T) {
	m, _ := press(t, newTestModel(t), tea.KeyPressMsg{Code: tea.KeyEnter})
	m, _ = press(t, m, key('b'))
	if m.choice != "" || m.content != "" {
		t.Errorf("choice=%q content=%q, want empty", m.choice, m.content)
	}
}

func TestQuit(t *testing.T) {
	for name, msg := range map[string]tea.Msg{
		"q":      key('q'),
		"ctrl+c": tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl},
	} {
		t.Run(name, func(t *testing.T) {
			m, cmd := press(t, newTestModel(t), msg)
			if !m.quitting {
				t.Error("quitting not set")
			}
			if cmd == nil {
				t.Fatal("expected quit command")
			}
			if _, ok := cmd().(tea.QuitMsg); !ok {
				t.Errorf("command did not produce QuitMsg")
			}
		})
	}
}

func TestViewListAndDetail(t *testing.T) {
	m := newTestModel(t)
	v := m.View()
	if !strings.Contains(v.Content, "RSS Reader") {
		t.Errorf("list view missing title: %q", v.Content)
	}
	if !v.AltScreen {
		t.Error("expected alt screen")
	}
	if v.MouseMode != tea.MouseModeCellMotion {
		t.Error("expected cell motion mouse mode")
	}

	m, _ = press(t, m, tea.KeyPressMsg{Code: tea.KeyEnter})
	if got := ansi.Strip(m.View().Content); !strings.Contains(got, "First episode") {
		t.Errorf("detail view missing episode title: %q", got)
	}
}
