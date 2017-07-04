package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var m *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func init() {
	/* load test data */
	setupFibNumbers()
	setupFibMap()
	setup()
}

func TestFibArrayBuild(t *testing.T) {

	if gFibNumbers[0] != 0 {
		fmt.Printf("fibonacci number at spot 4 is %d  and should be 2. \n", gFibNumbers[3])
		t.Fail()
	} else if gFibNumbers[93] != 12200160415121876738 {
		fmt.Printf("fibonacci number at spot 93 is %d  and should be 12200160415121876738. \n", gFibNumbers[3])
		t.Fail()
	}
}

func TestBuildFib(t *testing.T) {

	bfibResult, err := getFib(5)
	fibResult := string(bfibResult[:])

	if fibResult != "[0,1,1,2,3]" || err != nil {
		t.Fail()
	}
}

/* -- All the fibonacci numbers

[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,
10946,17711,28657,46368,75025,121393,196418,317811,514229,832040,1346269,
2178309,3524578,5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,
267914296,433494437,701408733,1134903170,1836311903,2971215073,4807526976,7778742049,12586269025,
20365011074,32951280099,53316291173,86267571272,139583862445,225851433717,365435296162,591286729879,
956722026041,1548008755920,2504730781961,4052739537881,6557470319842,10610209857723,17167680177565,
27777890035288,44945570212853,72723460248141,117669030460994,190392490709135,308061521170129,
498454011879264,806515533049393,1304969544928657,2111485077978050,3416454622906707,5527939700884757,
8944394323791464,14472334024676221,23416728348467685,37889062373143906,61305790721611591,
99194853094755497,160500643816367088,259695496911122585,420196140727489673,679891637638612258,
1100087778366101931,1779979416004714189,2880067194370816120,4660046610375530309,7540113804746346429,
12200160415121876738]
*/

func TestBuildFibMaxValue(t *testing.T) {

	allFibNumbers := "[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711," +
		"28657,46368,75025,121393,196418,317811,514229,832040,1346269,2178309,3524578," +
		"5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,267914296," +
		"433494437,701408733,1134903170,1836311903,2971215073,4807526976,7778742049," +
		"12586269025,20365011074,32951280099,53316291173,86267571272,139583862445," +
		"225851433717,365435296162,591286729879,956722026041,1548008755920,2504730781961," +
		"4052739537881,6557470319842,10610209857723,17167680177565,27777890035288," +
		"44945570212853,72723460248141,117669030460994,190392490709135,308061521170129," +
		"498454011879264,806515533049393,1304969544928657,2111485077978050,3416454622906707," +
		"5527939700884757,8944394323791464,14472334024676221,23416728348467685," +
		"37889062373143906,61305790721611591,99194853094755497,160500643816367088," +
		"259695496911122585,420196140727489673,679891637638612258,1100087778366101931," +
		"1779979416004714189,2880067194370816120,4660046610375530309,7540113804746346429," +
		"12200160415121876738]"

	bfibResult, err := getFib(18446744073709551615)
	fibResult := string(bfibResult[:])

	if fibResult != allFibNumbers || err != nil {
		t.Fail()
	}
}
func TestBuildFibMaxValueMinusOne(t *testing.T) {

	allFibNumbers := "[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711," +
		"28657,46368,75025,121393,196418,317811,514229,832040,1346269,2178309,3524578," +
		"5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,267914296," +
		"433494437,701408733,1134903170,1836311903,2971215073,4807526976,7778742049," +
		"12586269025,20365011074,32951280099,53316291173,86267571272,139583862445," +
		"225851433717,365435296162,591286729879,956722026041,1548008755920,2504730781961," +
		"4052739537881,6557470319842,10610209857723,17167680177565,27777890035288," +
		"44945570212853,72723460248141,117669030460994,190392490709135,308061521170129," +
		"498454011879264,806515533049393,1304969544928657,2111485077978050,3416454622906707," +
		"5527939700884757,8944394323791464,14472334024676221,23416728348467685," +
		"37889062373143906,61305790721611591,99194853094755497,160500643816367088," +
		"259695496911122585,420196140727489673,679891637638612258,1100087778366101931," +
		"1779979416004714189,2880067194370816120,4660046610375530309,7540113804746346429]"

	bfibResult, err := getFib(12200160415121876737)
	fibResult := string(bfibResult[:])

	if fibResult != allFibNumbers || err != nil {
		fmt.Printf("fib resutle = %s", fibResult)
		t.Fail()
	}
}

func TestBuildFibZero(t *testing.T) {

	bfibResult, err := getFib(0)
	fibResult := string(bfibResult[:])

	if fibResult != "[]" || err != nil {
		t.Fail()
	}
}
func TestBuildFibOne(t *testing.T) {

	bfibResult, err := getFib(1)
	fibResult := string(bfibResult[:])

	if fibResult != "[0]" || err != nil {
		t.Fail()
	}
}

func TestBuildFibBetween(t *testing.T) {

	bfibResult, err := getFib(4)
	fibResult := string(bfibResult[:])

	if fibResult != "[0,1,1,2,3]" || err != nil {
		t.Fail()
	}
}

func AddQuestionRoutes(r *mux.Router) {

	r.HandleFunc("/fibonacci", fibHandler).Methods("GET")
}

func setup() {
	//mux router with added question routes
	m = mux.NewRouter()
	AddQuestionRoutes(m)
	//The response recorder used to record HTTP responses
	respRec = httptest.NewRecorder()
}

func TestFibonacciFunctionalHandler(t *testing.T) {

	tests := []struct {
		description        string
		url                string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			description:        "simple fibonacci test",
			url:                "/fibonacci?fib=5",
			expectedStatusCode: 200,
			expectedBody:       "[0,1,1,2,3]",
		},
		{
			description:        "non fibonacci number input",
			url:                "/fibonacci?fib=4",
			expectedStatusCode: 200,
			expectedBody:       "[0,1,1,2,3]",
		},
		{
			description:        "fibonacci test zero",
			url:                "/fibonacci?fib=0",
			expectedStatusCode: 200,
			expectedBody:       "[]",
		},
		{
			description:        "fibonacci invalid parm",
			url:                "/fibonacci?fib=a",
			expectedStatusCode: 400,
			expectedBody:       "Invalid Input",
		},
		{
			description:        "fibonacci negative number",
			url:                "/fibonacci?fib=-1",
			expectedStatusCode: 400,
			expectedBody:       "Invalid Input",
		},
		{
			description:        "fibonacci out of bounds",
			url:                "/fibonacci?fib=18446744073709551616",
			expectedStatusCode: 400,
			expectedBody:       "Invalid Input",
		},
		{
			description:        "fibonacci floating point",
			url:                "/fibonacci?fib=1.0",
			expectedStatusCode: 400,
			expectedBody:       "Invalid Input",
		},
		{
			description:        "fibonacci empty input",
			url:                "/fibonacci?fib=",
			expectedStatusCode: 400,
			expectedBody:       "Invalid Input",
		},
		{
			description:        "fibonacci all fibonacci numbers ",
			url:                "/fibonacci?fib=18446744073709551615",
			expectedStatusCode: 200,
			expectedBody: "[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711," +
				"28657,46368,75025,121393,196418,317811,514229,832040,1346269,2178309,3524578," +
				"5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,267914296," +
				"433494437,701408733,1134903170,1836311903,2971215073,4807526976,7778742049," +
				"12586269025,20365011074,32951280099,53316291173,86267571272,139583862445," +
				"225851433717,365435296162,591286729879,956722026041,1548008755920,2504730781961," +
				"4052739537881,6557470319842,10610209857723,17167680177565,27777890035288," +
				"44945570212853,72723460248141,117669030460994,190392490709135,308061521170129," +
				"498454011879264,806515533049393,1304969544928657,2111485077978050,3416454622906707," +
				"5527939700884757,8944394323791464,14472334024676221,23416728348467685," +
				"37889062373143906,61305790721611591,99194853094755497,160500643816367088," +
				"259695496911122585,420196140727489673,679891637638612258,1100087778366101931," +
				"1779979416004714189,2880067194370816120,4660046610375530309,7540113804746346429,12200160415121876738]",
		},
		{
			description:        "fibonacci last fibonacci numbers ",
			url:                "/fibonacci?fib=12200160415121876738",
			expectedStatusCode: 200,
			expectedBody: "[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711," +
				"28657,46368,75025,121393,196418,317811,514229,832040,1346269,2178309,3524578," +
				"5702887,9227465,14930352,24157817,39088169,63245986,102334155,165580141,267914296," +
				"433494437,701408733,1134903170,1836311903,2971215073,4807526976,7778742049," +
				"12586269025,20365011074,32951280099,53316291173,86267571272,139583862445," +
				"225851433717,365435296162,591286729879,956722026041,1548008755920,2504730781961," +
				"4052739537881,6557470319842,10610209857723,17167680177565,27777890035288," +
				"44945570212853,72723460248141,117669030460994,190392490709135,308061521170129," +
				"498454011879264,806515533049393,1304969544928657,2111485077978050,3416454622906707," +
				"5527939700884757,8944394323791464,14472334024676221,23416728348467685," +
				"37889062373143906,61305790721611591,99194853094755497,160500643816367088," +
				"259695496911122585,420196140727489673,679891637638612258,1100087778366101931," +
				"1779979416004714189,2880067194370816120,4660046610375530309,7540113804746346429]",
		},
	}
	//-----------------------------------------------------------------
	// The below set of functional tests go through the handler.
	//-----------------------------------------------------------------
	for _, tc := range tests {

		req, err = http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Fatal("Creating 'GET /fibonacci' request failed!")
		}
		respRec = httptest.NewRecorder()

		m.ServeHTTP(respRec, req)

		if respRec.Code == tc.expectedStatusCode {
			if respRec.Body.String() == tc.expectedBody {
				t.Logf(tc.description + " - PASSED")
			} else {
				t.Fatal("Server error: Returned ", respRec.Body.String(), " instead of ", tc.expectedBody)
			}
		} else {
			t.Fatal("Server error: Returned ", respRec.Code, " instead of ", tc.expectedStatusCode)
		}
	}
}
