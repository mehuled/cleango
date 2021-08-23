package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_deleteFile(t *testing.T) {
	filename := "test.txt"
	_, err := os.Create(filename)
	assert.Nil(t, err)
	err = os.Remove(filename)
	assert.Nil(t, err)
	exists , err := os.Stat(filename);
	fmt.Println(exists)

}
