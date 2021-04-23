package main

import (
	"context"
	"fmt"
)

func main() {
	//  Case WithTimeout
	/*
		ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
		ch := time.After(600 * time.Millisecond)

		select {
		case <-ctx.Done():
			log.Println("ctx.Done()")
			log.Printf("ctx.Err() Done : %q", ctx.Err())
		case <-ch:
			log.Println("Ch done")
			log.Printf("ctx.Err() Ch : %q", ctx.Err())
		}
	*/

	// Case WithCancel
	/*
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ch := time.After(600 * time.Millisecond)

		time.AfterFunc(400*time.Millisecond, func() {
			log.Println("We are about to cancel the ctx.")
			cancel()
		})
		select {
		case <-ctx.Done():
			log.Println("ctx.Done()")
			log.Printf("ctx.Err() Done : %q", ctx.Err())
		case <-ch:
			log.Println("Ch done")
			log.Printf("ctx.Err() Ch : %q", ctx.Err())
		}
	*/

	// Case Context With Value
	ctx, _ := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, "b", "1")
	ctx3 := context.WithValue(ctx2, "b", "2")
	ctx4 := context.WithValue(ctx3, "c", "3")

	lookup(ctx, "ctx", "a")
	lookup(ctx2, "ctx2", "b")
	lookup(ctx3, "ctx3", "a")
	lookup(ctx3, "ctx3", "b")
	lookup(ctx4, "ctx4", "b")
	lookup(ctx4, "ctx4", "b")

}

func lookup(ctx context.Context, name, k string) {
	fmt.Println(name, ctx.Value(k))

}
