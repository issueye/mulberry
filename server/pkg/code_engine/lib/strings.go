package lib

import (
	"mulberry/pkg/code_engine/lib/stand"
	original_strings "strings"
)

var (
	stringsModule = stand.NewGojaModule("go/strings")
)

func InitStrings() {
	stringsModule.Set(
		stand.Objects{
			// Functions
			"clone":          original_strings.Clone,
			"compare":        original_strings.Compare,
			"contains":       original_strings.Contains,
			"containsAny":    original_strings.ContainsAny,
			"containsFunc":   original_strings.ContainsFunc,
			"containsRune":   original_strings.ContainsRune,
			"count":          original_strings.Count,
			"cut":            original_strings.Cut,
			"cutPrefix":      original_strings.CutPrefix,
			"cutSuffix":      original_strings.CutSuffix,
			"equalFold":      original_strings.EqualFold,
			"fields":         original_strings.Fields,
			"fieldsFunc":     original_strings.FieldsFunc,
			"hasPrefix":      original_strings.HasPrefix,
			"hasSuffix":      original_strings.HasSuffix,
			"index":          original_strings.Index,
			"indexAny":       original_strings.IndexAny,
			"indexByte":      original_strings.IndexByte,
			"indexFunc":      original_strings.IndexFunc,
			"indexRune":      original_strings.IndexRune,
			"join":           original_strings.Join,
			"lastIndex":      original_strings.LastIndex,
			"lastIndexAny":   original_strings.LastIndexAny,
			"lastIndexByte":  original_strings.LastIndexByte,
			"lastIndexFunc":  original_strings.LastIndexFunc,
			"map":            original_strings.Map,
			"newReader":      original_strings.NewReader,
			"newReplacer":    original_strings.NewReplacer,
			"repeat":         original_strings.Repeat,
			"replace":        original_strings.Replace,
			"replaceAll":     original_strings.ReplaceAll,
			"split":          original_strings.Split,
			"splitAfter":     original_strings.SplitAfter,
			"splitAfterN":    original_strings.SplitAfterN,
			"splitN":         original_strings.SplitN,
			"toLower":        original_strings.ToLower,
			"toLowerSpecial": original_strings.ToLowerSpecial,
			"toTitle":        original_strings.ToTitle,
			"toTitleSpecial": original_strings.ToTitleSpecial,
			"toUpper":        original_strings.ToUpper,
			"toUpperSpecial": original_strings.ToUpperSpecial,
			"toValidUTF8":    original_strings.ToValidUTF8,
			"trim":           original_strings.Trim,
			"trimFunc":       original_strings.TrimFunc,
			"trimLeft":       original_strings.TrimLeft,
			"trimLeftFunc":   original_strings.TrimLeftFunc,
			"trimPrefix":     original_strings.TrimPrefix,
			"trimRight":      original_strings.TrimRight,
			"trimRightFunc":  original_strings.TrimRightFunc,
			"trimSpace":      original_strings.TrimSpace,
			"trimSuffix":     original_strings.TrimSuffix,

			// Var and consts

			// Types (value type)
			"Builder":  func() original_strings.Builder { return original_strings.Builder{} },
			"Reader":   func() original_strings.Reader { return original_strings.Reader{} },
			"Replacer": func() original_strings.Replacer { return original_strings.Replacer{} },

			// Types (pointer type)
			"NewBuilder":  func() *original_strings.Builder { return &original_strings.Builder{} },
			"NewReader":   func() *original_strings.Reader { return &original_strings.Reader{} },
			"NewReplacer": func() *original_strings.Replacer { return &original_strings.Replacer{} },
		},
	).Register()
}
