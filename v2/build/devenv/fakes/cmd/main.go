package main

import (
	"github.com/gin-gonic/gin"

	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/fake"
)

func main() {
	_, err := fake.NewFakeDataProvider(&fake.Input{Port: fake.DefaultFakeServicePort})
	err = fake.Func("POST", "/my_fake_api", func(ctx *gin.Context) {
		// TODO: add CCV specific mocks
		ctx.JSON(200, gin.H{
			"data": map[string]any{
				"result": "something",
			},
		})
	})
	if err != nil {
		panic(err)
	}
	select {}
}
