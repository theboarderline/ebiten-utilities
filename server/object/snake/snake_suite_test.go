package snake_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnake(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Snake Suite")
}
