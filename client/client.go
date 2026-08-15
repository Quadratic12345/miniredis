package main
import ("bufio";"fmt";"net";"os")
func main(){
	conn, err :=net.Dial("tcp","127.0.0.1:9999")
	if err !=nil{
		fmt.Println("connection error",err)
		return 
	}
	defer conn.Close()
	fmt.Println("connected to mini Redis")
	fmt.Println("commands: SET,GET,DEL")
	fmt.Println("Type exit to quit")
	input :=bufio.NewScanner(os.Stdin)
	response:=bufio.NewScanner(conn)
	for{
		fmt.Print(">")
		if !input.Scan(){
			break
		}
		command:=input.Text()
		if command=="exit"{
			break
		}
		fmt.Fprintln(conn,command)
		if response.Scan(){
			fmt.Println(response.Text())
		}
	}
}