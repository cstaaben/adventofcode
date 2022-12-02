package convert

import "strings"

func KebabToPascal(s string) string {
    res := ""

    for _, w := range strings.Split(s, "-") {
        res += strings.ToUpper(string(w[0])) + w[1:]
    }

    return res
}

func KebabToLower(s string) (lower string) {
    for _, w := range strings.Split(s, "-") {
        lower += strings.ToLower(w)
    }

    return lower
}

func KebabToTitle(s string) (title string) {
    for _, w := range strings.Split(s, "-") {
        title += strings.ToUpper(string(w[0])) + w[1:] + "-"
    }

    return strings.TrimSuffix(title, "-")
}

func KebabToCamel(s string) (camel string) {
    for i, w := range strings.Split(s, "-") {
        if i == 0 {
            camel += w
            continue
        }

        camel += strings.ToUpper(string(w[0])) + w[1:]
    }

    return camel
}

func CamelToPascal(s string) string {
    return strings.ToUpper(string(s[0])) + s[1:]
}
