package diff

import (
	"testing"

	"github.com/pirmd/verify"
)

func TestGetSameHead(t *testing.T) {
	testCases := []struct {
		l, r         []string
		wantl, wantr []string
		wantd        Result
	}{
		{[]string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, []string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, nil},
		{[]string{"a", "x", "c"}, []string{"a", "b", "c"}, []string{"x", "c"}, []string{"b", "c"}, Result{&diff{IsSame, "a"}}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}, []string{}, Result{&diff{IsSame, "a"}, &diff{IsSame, "b"}, &diff{IsSame, "c"}}},
		{[]string{"a", "b", "c"}, []string{"a", "b"}, []string{"c"}, []string{}, Result{&diff{IsSame, "a"}, &diff{IsSame, "b"}}},
	}

	for _, tc := range testCases {
		gotl, gotr, gotd := getSameHead(tc.l, tc.r)
		verify.Equal(t, gotl, tc.wantl, "Extracting same head failed for remaining l.")
		verify.Equal(t, gotr, tc.wantr, "Extracting same head failed for remaining r.")
		verify.Equal(t, gotd, tc.wantd, "Extracting same head failed for diff d.")
	}
}

func TestGetSameTail(t *testing.T) {
	testCases := []struct {
		l, r         []string
		wantl, wantr []string
		wantd        Result
	}{
		{[]string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, []string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, nil},
		{[]string{"a", "x", "c"}, []string{"a", "b", "c"}, []string{"a", "x"}, []string{"a", "b"}, Result{&diff{IsSame, "c"}}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{}, []string{}, Result{&diff{IsSame, "a"}, &diff{IsSame, "b"}, &diff{IsSame, "c"}}},
		{[]string{"a", "b", "c"}, []string{"b", "c"}, []string{"a"}, []string{}, Result{&diff{IsSame, "b"}, &diff{IsSame, "c"}}},
	}

	for _, tc := range testCases {
		gotl, gotr, gotd := getSameTail(tc.l, tc.r)
		verify.Equal(t, gotl, tc.wantl, "Extracting same tail failed for remaining l.")
		verify.Equal(t, gotr, tc.wantr, "Extracting same tail failed for remaining r.")
		verify.Equal(t, gotd, tc.wantd, "Extracting same tail failed for diff d.")
	}
}

func TestMatrixLCS(t *testing.T) {
	testCases := []struct {
		a, b []string
		want [][]int
	}{
		{[]string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 1, 1, 1, 1}, {0, 1, 1, 1, 2, 2}, {0, 1, 1, 2, 2, 2}}},
		{[]string{"a", "x", "c"}, []string{"a", "b", "c"}, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}, {0, 1, 1, 1}, {0, 1, 1, 2}}},
	}

	for _, tc := range testCases {
		got := matrixLCS(tc.a, tc.b)
		verify.Equal(t, got, tc.want, "LCS Matrix is incorrect:")
	}
}

func TestSequenceLCS(t *testing.T) {
	testCases := []struct {
		a, b []string
		want [][2]int
	}{
		{[]string{"g", "a", "c"}, []string{"a", "g", "c", "a", "t"}, [][2]int{{0, 1}, {1, 3}}},
		{[]string{"a", "x", "c"}, []string{"a", "b", "c"}, [][2]int{{0, 0}, {2, 2}}},
	}

	for _, tc := range testCases {
		got := sequenceLCS(tc.a, tc.b, false)
		verify.Equal(t, got, tc.want, "LCS sequence is incorrect")
	}
}

func TestVanillaLCS(t *testing.T) {
	testCases := []struct {
		l, r string
		want Result
	}{
		{
			l: `#include <stdio.h>

// Frobs foo heartily
int frobnitz(int foo)
{
    int i;
    for(i = 0; i < 10; i++)
    {
        printf("Your answer is: ");
        printf("%d\n", foo);
    }
}

int fact(int n)
{
    if(n > 1)
    {
        return fact(n-1) * n;
    }
    return 1;
}

int main(int argc, char **argv)
{
    frobnitz(fact(10));
}`,

			r: `#include <stdio.h>

int fib(int n)
{
    if(n > 2)
    {
        return fib(n-1) + fib(n-2);
    }
    return 1;
}

// Frobs foo heartily
int frobnitz(int foo)
{
    int i;
    for(i = 0; i < 10; i++)
    {
        printf("%d\n", foo);
    }
}

int main(int argc, char **argv)
{
    frobnitz(fib(10));
}`,

			want: Result{
				&diff{IsSame, "#include <stdio.h>\n"},
				&diff{IsSame, "\n"},
				&diff{IsInserted, "int fib(int n)\n"},
				&diff{IsInserted, "{\n"},
				&diff{IsInserted, "    if(n > 2)\n"},
				&diff{IsInserted, "    {\n"},
				&diff{IsInserted, "        return fib(n-1) + fib(n-2);\n"},
				&diff{IsInserted, "    }\n"},
				&diff{IsInserted, "    return 1;\n"},
				&diff{IsInserted, "}\n"},
				&diff{IsInserted, "\n"},
				&diff{IsSame, "// Frobs foo heartily\n"},
				&diff{IsSame, "int frobnitz(int foo)\n"},
				&diff{IsSame, "{\n"},
				&diff{IsSame, "    int i;\n"},
				&diff{IsSame, "    for(i = 0; i < 10; i++)\n"},
				&diff{IsSame, "    {\n"},
				&diff{IsDeleted, "        printf(\"Your answer is: \");\n"},
				&diff{IsSame, "        printf(\"%d\\n\", foo);\n"},
				&diff{IsSame, "    }\n"},
				&diff{IsSame, "}\n"},
				&diff{IsSame, "\n"},
				&diff{IsDeleted, "int fact(int n)\n"},
				&diff{IsDeleted, "{\n"},
				&diff{IsDeleted, "    if(n > 1)\n"},
				&diff{IsDeleted, "    {\n"},
				&diff{IsDeleted, "        return fact(n-1) * n;\n"},
				&diff{IsDeleted, "    }\n"},
				&diff{IsDeleted, "    return 1;\n"},
				&diff{IsDeleted, "}\n"},
				&diff{IsDeleted, "\n"},
				&diff{IsSame, "int main(int argc, char **argv)\n"},
				&diff{IsSame, "{\n"},
				&diff{IsDeleted, "    frobnitz(fact(10));\n"},
				&diff{IsInserted, "    frobnitz(fib(10));\n"},
				&diff{IsSame, "}\n"},
			},
		},
	}

	for _, tc := range testCases {
		got := VanillaLCS(ByLines(tc.l), ByLines(tc.r))
		verify.Equal(t, got, tc.want, "Vanilla LCS diff is not as expected.")
	}
}
