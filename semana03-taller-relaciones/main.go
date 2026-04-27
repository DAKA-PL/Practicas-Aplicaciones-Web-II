package main

import (
	"fmt"
	"semana03-taller-relaciones/internal/cafeteria"
)

func main() {

	var repo cafeteria.Repository = cafeteria.NewRepoMemoria()

	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Ana"})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Pepe"})

	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Agua", Precio: 1, Stock: 60, Categoria: "Bebida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Pan", Precio: 2, Stock: 40, Categoria: "Comida"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Cafe", Precio: 3, Stock: 90, Categoria: "Bebida"})

	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Printf("Error al buscar cliente: %s\n", err.Error())
	} else {
		fmt.Printf("\nEncontrado: %s\n", c.Nombre)
	}

	fantasma, err := repo.ObtenerCliente(3)
	if err != nil {
		fmt.Printf("Error al buscar cliente: %s\n", err.Error())
	} else {
		fmt.Println(fantasma)
	}

	fmt.Printf("Listando Productos:\n")
	for _, p := range repo.ListarProductos() {
		fmt.Printf("%v\n", p)
	}
}

/*
   	Daivelyn Pincay Lopez
   	1. ¿Tuviste que poner Cliente, Producto y Pedido en el mismo paquete? ¿Por qué sí o por qué no?
   	Si se tuvo que poner,porque en caso contrario de no colocarlo se tendria que crear otro paquete

   	2. ¿Qué problema aparecería si intentaras separar Producto en un paquete aparte cuando Pedido lotiene anidado?
       El problema seria que al momento de ejecutarlo daria error ya que no esta llamando a las funciones

   	3. Comparando con el Día A (donde usamos IDs): ¿qué ventaja tiene el modelo con IDs paraorganizar el código en paquetes?
   	La ventaja que tiene es que los errores serian facil de detectar a tiempo

   	Nahim Simba Mero
   	1. ¿Tuviste que poner Cliente, Producto y Pedido en el mismo paquete? ¿Por qué sí o por qué no?
   	Si ya que voy a trabajar con ellos en el paquete cafeteria, sino tendria que crear otro paquete mas

   	2. ¿Qué problema aparecería si intentaras separar Producto en un paquete aparte cuando Pedido lo tiene anidado?
   	Tendria que llamarlo a mi paquete cafeteria para poder usarlo en las diferentes funciones ademas de corregir el main

   	3. Comparando con el Día A (donde usamos IDs): ¿qué ventaja tiene el modelo con IDs paraorganizar el código en paquetes?
   	Permite ser mas escalable y permite una mayor eficiencia en mantenimientos

*/
