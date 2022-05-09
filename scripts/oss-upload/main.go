//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/KusionStack/kusionup/pkg/client/oss"
)

const (
	endPoint                     = "TODO"
	bucketName                   = "kusion-public"
	kusionDarwinBucketKey        = "cli/kusionup/darwin/bin/kusionup"
	kusionDarwinArm64BucketKey   = "cli/kusionup/darwin-arm64/bin/kusionup"
	kusionLinuxBucketKey         = "cli/kusionup/linux/bin/kusionup"
	kusionWindowsBucketKey       = "cli/kusionup/windows/bin/kusionup.exe"
	kusionDarwinReleasePath      = "build/darwin/bin/kusionup"
	kusionDarwinArm64ReleasePath = "build/darwin-arm64/bin/kusionup"
	kusionLinuxReleasePath       = "build/linux/bin/kusionup"
	kusionWindowsReleasePath     = "build/windows/bin/kusionup.exe"
)

func main() {
	id := os.Getenv("OSS_ACCESS_KEY_ID")
	secret := os.Getenv("OSS_ACCESS_KEY_SECRET")
	client, err := oss.NewRemoteClient(endPoint, id, secret, "", bucketName)
	if err != nil {
		panic(err)
	}
	if fileExists(kusionDarwinReleasePath) {
		err = client.PutObjectFromFile(kusionDarwinBucketKey, kusionDarwinReleasePath)
		if err != nil {
			panic(err)
		}
		if fileExists(kusionDarwinReleasePath + ".md5.txt") {
			err = client.PutObjectFromFile(kusionDarwinBucketKey+".md5.txt", kusionDarwinReleasePath+".md5.txt")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Upload darwin artifact to OSS Successfully!")
	}
	if fileExists(kusionDarwinArm64ReleasePath) {
		err = client.PutObjectFromFile(kusionDarwinArm64BucketKey, kusionDarwinArm64ReleasePath)
		if err != nil {
			panic(err)
		}
		if fileExists(kusionDarwinArm64ReleasePath + ".md5.txt") {
			err = client.PutObjectFromFile(kusionDarwinArm64BucketKey+".md5.txt", kusionDarwinArm64ReleasePath+".md5.txt")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Upload darwin-arm64 artifact to OSS Successfully!")
	}
	if fileExists(kusionLinuxReleasePath) {
		err = client.PutObjectFromFile(kusionLinuxBucketKey, kusionLinuxReleasePath)
		if err != nil {
			panic(err)
		}
		if fileExists(kusionLinuxReleasePath + ".md5.txt") {
			err = client.PutObjectFromFile(kusionLinuxBucketKey+".md5.txt", kusionLinuxReleasePath+".md5.txt")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Upload linux artifact to OSS Successfully!")
	}
	if fileExists(kusionWindowsReleasePath) {
		err = client.PutObjectFromFile(kusionWindowsBucketKey, kusionWindowsReleasePath)
		if err != nil {
			panic(err)
		}
		if fileExists(kusionWindowsReleasePath + ".md5.txt") {
			err = client.PutObjectFromFile(kusionWindowsBucketKey+".md5.txt", kusionWindowsReleasePath+".md5.txt")
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("Upload windows artifact to OSS Successfully!")
	}
}

func fileExists(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	}
	return true
}
