package main

    import (
        "errors"
        "fmt"
		"log"
		"net"
        "net/http"
        "net/rpc"
		"os"
    )

    type Args struct {
        A, B int
    }

    type Quotient struct {
        Quo, Rem int
    }

    type Arith int

    func (t *Arith) Multiply(args *Args, reply *int) error {
        *reply = args.A * args.B
        return nil
    }

    func (t *Arith) Divide(args *Args, quo *Quotient) error {
        if args.B == 0 {
            return errors.New("divide by zero")
        }
        quo.Quo = args.A / args.B
        quo.Rem = args.A % args.B
        return nil
    }

    func main() {
		
        arith := new(Arith)
        rpc.Register(arith)
        rpc.HandleHTTP()

        lis, err := net.Listen("tcp", "127.0.0.1:8095")
        if err != nil {
            fmt.Println(err.Error())
        }
		fmt.Printf(os.Sdout, "%s", "start connection")
		http.Serve(lis, nil)
    }