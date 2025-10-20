package main

import (
	"fmt"

	"github.com/AppBlitz/strutureData/lists"
)

func main() {
	fmt.Printf("%s\n", "Lista enlazada simple")
	singlyLinkedList := lists.SinglyLinkedList{}
	fmt.Printf("%s   %t\n", "Miramos si la lista se encuentra vacia", singlyLinkedList.IsEmpty())
	singlyLinkedList.AppendEnd(1)
	singlyLinkedList.AppendEnd(2)
	singlyLinkedList.AppendEnd(3)
	singlyLinkedList.AppendStart(0)
	singlyLinkedList.AppendStart(-1)
	fmt.Printf("%s  %d\n", "Longitud de la lista ", singlyLinkedList.Len())
	singlyLinkedList.RemoveElementPosition(3)
	fmt.Printf("%s  %d\n", "Longitud de la lista después de eliminar un elemento ", singlyLinkedList.Len())
	singlyLinkedList.Display()
	singlyLinkedList.RemoveAll()
	fmt.Printf("%s   %t\n", "Miramos si la lista se encuentra vacia", singlyLinkedList.IsEmpty())

	fmt.Printf("%s\n", "Pilas")
	stack := lists.Stack{}
	fmt.Printf("%s   %t\n", "Miramos si la pila se encuentra vacia", stack.StackEmpty())
	stack.Push(-1)
	stack.Push(0)
	stack.Push(1)
	stack.Push(2)
	stack.Display()
	stack.Pop()
	stack.Display()

	fmt.Printf("%s\n", "Lista circular")
	circulyLinkedList := lists.LinkedListCircular{}
	fmt.Printf("%s   %t\n", "Miramos si la lista enlazada doble esta vacia", circulyLinkedList.IsEmpty())
	circulyLinkedList.AppendEnd(1)
	circulyLinkedList.AppendEnd(2)
	circulyLinkedList.AppendEnd(3)
	circulyLinkedList.AppendEnd(4)
	circulyLinkedList.AppendEnd(5)
	circulyLinkedList.AppendEnd(6)
	circulyLinkedList.AppendStart(-1)
	circulyLinkedList.AppendStart(-3)
	circulyLinkedList.AppendStart(-2)
	fmt.Printf("%s   %t\n", "Miramos si la lista enlazada doble esta vacia", circulyLinkedList.IsEmpty())
	circulyLinkedList.RemovePoisition(4)
	circulyLinkedList.InsertInpositionRecursive(5, 8)
	circulyLinkedList.RemoveAll()
	fmt.Printf("%s   %t\n", "Miramos si la lista enlazada doble esta vacia", circulyLinkedList.IsEmpty())

	fmt.Printf("%s\n", "Cola")
	queue := lists.Queue{}
	fmt.Printf("%s   %d\n", "Miramos la longitud de la cola", queue.Len())
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.EnQueue(6)
	queue.EnQueue(7)
	queue.EnQueue(5)
	queue.Display()

	fmt.Printf("%s\n", "Lista enlazada boblemente enlazada")
	circulyDoublyLinkedList := lists.CircularDoublyLinke{}
	fmt.Printf("%s   %t", "Verificamos si la lista doblemente enlazada esta vacia", circulyDoublyLinkedList.IsEmpty())
	circulyDoublyLinkedList.AppendEnd(2)
	circulyDoublyLinkedList.AppendEnd(3)
	circulyDoublyLinkedList.AppendEnd(4)
	circulyDoublyLinkedList.AppendEnd(5)
	circulyDoublyLinkedList.AppendEnd(6)
	circulyDoublyLinkedList.AppendEnd(7)
	circulyDoublyLinkedList.AppendEnd(8)
	circulyDoublyLinkedList.AppendStart(1)
	circulyDoublyLinkedList.AppendStart(0)
	circulyDoublyLinkedList.AppendStart(-1)
	circulyDoublyLinkedList.AppendStart(-2)
	circulyDoublyLinkedList.AppendStart(-3)
	circulyDoublyLinkedList.Display()
	circulyDoublyLinkedList.InsertarEnPosicionRecursivo(3, 5)
	circulyDoublyLinkedList.RemoveAt(6)
	fmt.Printf("%s   %t\n", "Verificamos si la lista doblemente enlazada esta vacia", circulyDoublyLinkedList.IsEmpty())
	fmt.Printf("%s %t\n", "Buscamos un dato recursivo", circulyDoublyLinkedList.BuscarRecursivo(-3))
	circulyDoublyLinkedList.Display()

	circulyDoublyLinkedList.RemoveAll()
	fmt.Printf("%s   %t\n", "Verificamos si la lista doblemente enlazada esta vacia", circulyDoublyLinkedList.IsEmpty())

	tree := lists.NewTree()

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	tree.PrintInorder() // 20 30 40 50 60 70 80
	tree.Delete(70)
	tree.PrintInorder() // 20 30 40 50 60 80

	fmt.Println("Height:", tree.Height())
	fmt.Println("Nodes:", tree.CountNodes())
	fmt.Println("Sum:", tree.Sum())
}
