package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"project/commit_info"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func main() {

	var conn *grpc.ClientConn
	host := os.Getenv("INPUT_HOST")
	conn, err := grpc.Dial(host, grpc.WithInsecure())

	CheckErr(err, "did not connect: %s")

	defer conn.Close()

	c := commit_info.NewCommitDataClient(conn)

	masterHeadCommit := os.Getenv("INPUT_HEADHASH")
	diff := os.Getenv("INPUT_DIFF")
	command := os.Getenv("INPUT_COMMAND")

	response, err := c.Translate(context.Background(), &commit_info.CommitInfo{
		HeadHash: masterHeadCommit, CommitDiff: diff, CommandLine: command})
	CheckErr(err, "Error when processing git info: %s")
	log.Printf("Response from server: %s", response)

}
