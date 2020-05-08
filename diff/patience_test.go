package diff

import (
	"reflect"
	"testing"
)

func TestVanillaPatience(t *testing.T) {
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
				&diff{IsSame, "}"},
			},
		},
	}

	for _, tc := range testCases {
		got := VanillaPatience(ByLines(tc.l), ByLines(tc.r))
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Patience diff between %v and %v failed.\nWant: %v\nGot : %v.", tc.l, tc.r, tc.want, got)
		}
	}
}
