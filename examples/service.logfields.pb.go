// Code generated by protoc-gen-gologfields. DO NOT EDIT.

package example

import (
	fmt "fmt"
	go_proto_logfields "github.com/nvx/go-proto-logfields"
)

func (m *Note) LogFields() map[string]string {
	// Handle being called on nil message.
	if m == nil {
		return map[string]string{}
	}

	literalsMap := map[string]string{
		"author": m.GetAuthor(),
	}

	return literalsMap
}

func (m *Note) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if m == nil {
		return
	}

	dst["author"] = m.GetAuthor()
}

func (m *Request) LogFields() map[string]string {
	// Handle being called on nil message.
	if m == nil {
		return map[string]string{}
	}

	// Gather fields from oneofs and child messages.
	var hasInner bool

	childMsgMapNote := go_proto_logfields.ExtractLogFieldsFromMessage(m.Note)
	hasInner = hasInner || len(childMsgMapNote) > 0

	literalsMap := map[string]string{
		"path": m.GetPath(),
	}

	if !hasInner {
		return literalsMap
	}

	for k, v := range childMsgMapNote {
		literalsMap[k] = v
	}

	return literalsMap
}

func (m *Request) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if m == nil {
		return
	}

	dst["path"] = m.GetPath()

	go_proto_logfields.ExtractRequestFieldsFromMessage(m.GetNote(), dst)
}

func (m *Response) LogFields() map[string]string {
	// Handle being called on nil message.
	if m == nil {
		return map[string]string{}
	}

	// Gather fields from oneofs and child messages.
	var hasInner bool

	childMsgMapChangedNote := go_proto_logfields.ExtractLogFieldsFromMessage(m.ChangedNote)
	hasInner = hasInner || len(childMsgMapChangedNote) > 0

	literalsMap := map[string]string{
		"did_it": fmt.Sprintf("%v", m.GetDidStuff()),
	}

	if !hasInner {
		return literalsMap
	}

	for k, v := range childMsgMapChangedNote {
		literalsMap[k] = v
	}

	return literalsMap
}

func (m *Response) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if m == nil {
		return
	}

	dst["did_it"] = m.GetDidStuff()

	go_proto_logfields.ExtractRequestFieldsFromMessage(m.GetChangedNote(), dst)
}
