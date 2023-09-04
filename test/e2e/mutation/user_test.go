package mutation_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/pkg/infrastructure/router"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/testutil"
	"github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/testutil/e2e"
	"github.com/gavv/httpexpect/v2"
)

func TestUser_CreateUser(t *testing.T) {
	expect, client, teardown := e2e.Setup(t, e2e.SetupOption{
		Teardown: func(t *testing.T, client *ent.Client) {
			testutil.DropUser(t, client)
		},
	})
	defer teardown()

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(t *testing.T) *httpexpect.Response
		assert  func(t *testing.T, got *httpexpect.Response)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name:    "It should create test user",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {
				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {
							createUser(input: {name: "Tom1", age: 20}) {
								age
								name
								id
								createdAt
								updatedAt
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got).Object()
				testUser := e2e.GetObject(data, "createUser")
				testUser.Value("age").Number().IsEqual(20)
				testUser.Value("name").String().IsEqual("Tom1")
			},
			teardown: func(t *testing.T) {
				testutil.DropUser(t, client)
			},
		},
		{
			name:    "It should NOT create test user when the length of the name is over",
			arrange: func(t *testing.T) {},
			act: func(t *testing.T) *httpexpect.Response {
				return expect.POST(router.QueryPath).WithJSON(map[string]string{
					"query": `
						mutation {  
							createUser(input: {name: "Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1Tom1", age: 20}) {   
								age    
								name
								id    
								createdAt    
								updatedAt  
						}
					}`,
				}).Expect()
			},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.Status(http.StatusOK)
				data := e2e.GetData(got)
				data.IsNull()

				errors := e2e.GetErrors(got)
				errors.Array().Length().IsEqual(1)
			},
			teardown: func(t *testing.T) {
				testutil.DropUser(t, client)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got := tt.act(t)
			tt.assert(t, got)
			tt.teardown(t)
		})
	}
}
