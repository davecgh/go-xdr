/*
 * Copyright (c) 2015-2017 Giovanni Bajo <rasky@develer.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package xdr

import (
	"reflect"
	"testing"
)

type structWithTag struct {
	A int `xdr:"a=123,b,c=true,d=false"`
}

func TestTag(t *testing.T) {
	s := structWithTag{}
	rt := reflect.TypeOf(s).Field(0)
	tag := parseTag(rt.Tag)

	if tag.Get("a") != "123" {
		t.Errorf("wrong value for a: %v", tag.Get("a"))
	}
	if tag.Get("b") != "true" {
		t.Errorf("wrong value for b: %v", tag.Get("b"))
	}
	if tag.Get("c") != "true" {
		t.Errorf("wrong value for b: %v", tag.Get("c"))
	}
	if tag.Get("d") != "false" {
		t.Errorf("wrong value for b: %v", tag.Get("d"))
	}
}
