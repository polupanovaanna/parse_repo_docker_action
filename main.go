package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"project/util"
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

	c := util.NewCommitDataClient(conn)

	masterHeadCommit := os.Getenv("INPUT_HEADHASH")
	diff := os.Getenv("INPUT_DIFF")

	response, err := c.Translate(context.Background(), &util.CommitInfo{HeadHash: masterHeadCommit, CommitDiff: diff})
	CheckErr(err, "Error when processing git info: %s")
	log.Printf("Response from server: %s", response)

}
