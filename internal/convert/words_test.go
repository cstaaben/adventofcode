package convert

import "testing"

func TestKebabToPascal(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name     string
        word     string
        expected string
    }{
        {
            name:     "One word",
            word:     "one",
            expected: "One",
        },
        {
            name:     "Two words",
            word:     "forty-two",
            expected: "FortyTwo",
        },
        {
            name:     "Five words",
            word:     "how-now-brown-cow-abc",
            expected: "HowNowBrownCowAbc",
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if w := KebabToPascal(tc.word); w != tc.expected {
                t.Errorf("expected %s but received %s", tc.expected, w)
            }
        })
    }
}

func TestKebabToLower(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name     string
        word     string
        expected string
    }{
        {
            name:     "one word",
            word:     "one",
            expected: "one",
        },
        {
            name:     "two words",
            word:     "forty-two",
            expected: "fortytwo",
        },
        {
            name:     "four words",
            word:     "how-now-brown-cow",
            expected: "hownowbrowncow",
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if w := KebabToLower(tc.word); w != tc.expected {
                t.Errorf("expected %s but received %s", tc.expected, w)
            }
        })
    }
}

func TestKebabToTitle(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name     string
        word     string
        expected string
    }{
        {
            name:     "one word",
            word:     "one",
            expected: "One",
        },
        {
            name:     "two words",
            word:     "forty-two",
            expected: "Forty-Two",
        },
        {
            name:     "four words",
            word:     "how-now-brown-cow",
            expected: "How-Now-Brown-Cow",
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if w := KebabToTitle(tc.word); w != tc.expected {
                t.Errorf("expected %s but received %s", tc.expected, w)
            }
        })
    }
}

func TestKebabToCamel(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name     string
        word     string
        expected string
    }{
        {
            name:     "one word",
            word:     "one",
            expected: "one",
        },
        {
            name:     "two words",
            word:     "forty-two",
            expected: "fortyTwo",
        },
        {
            name:     "four words",
            word:     "how-now-brown-cow",
            expected: "howNowBrownCow",
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if w := KebabToCamel(tc.word); w != tc.expected {
                t.Errorf("expected %s but received %s", tc.expected, w)
            }
        })
    }
}

func TestCamelToPascal(t *testing.T) {
    t.Parallel()

    testCases := []struct {
        name     string
        word     string
        expected string
    }{
        {
            name:     "one word",
            word:     "one",
            expected: "One",
        },
        {
            name:     "two words",
            word:     "fortyTwo",
            expected: "FortyTwo",
        },
        {
            name:     "four words",
            word:     "howNowBrownCow",
            expected: "HowNowBrownCow",
        },
    }

    for _, tc := range testCases {
        tc := tc
        t.Run(tc.name, func(t *testing.T) {
            if w := CamelToPascal(tc.word); w != tc.expected {
                t.Errorf("expected %s but received %s", tc.expected, w)
            }
        })
    }
}
