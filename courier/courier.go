package courier

/*
Recordar, este archivo guarda la implementación de las funciones que se utilizarán
junto a las estructuras de datos que necesitamos en server.go
*/

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/context"
)

//Server Se declara la estructura del servidor
type Server struct {
	//Mapa para guardar los pares "CodigoSeguimiento-Estado" de cada paquete
	MapaSeguimiento map[int64]string

	//Slice que tendrá las órdenes de pymes que lleguen al servidor
	OrdenesP []*OrdenPyme

	//Slice que tendrá las órdens de retail que lleguen al servidor
	OrdenesR []*OrdenRetail
	//mutex que protegerá las variables compartidas
	MuMapaSeg sync.Mutex
}

//Seguimiento Esta es la prueba inicial de mi función para devolver código de seguimiento
func (s *Server) Seguimiento(ctx context.Context, seguim *Codigo) (*Mensaje, error) {
	//Buscamos el codigo en el mapa de seguimiento
	i, ok := s.MapaSeguimiento[seguim.Cod]
	if !ok {
		err := errors.New("Error: missing key in MapaSeguimiento")
		return &Mensaje{Body: "Ese codigo no está registrado"}, err
	}
	return &Mensaje{Body: i}, nil
}

//EnviarOrdenPyme se recibe una orden, se guarda en el slice "ordenes"
//
func (s *Server) EnviarOrdenPyme(ctx context.Context, ordenP *OrdenPyme) (*Codigo, error) {

	//Implementacion de una fuente de numeros aleatorios
	//para dar códigos diferentes en cada instancia
	seed := rand.NewSource(time.Now().UnixNano())
	source := rand.New(seed)

	var num int64
	var err error
	for {
		// Se usa un set de numeros (Cardinalidad 1x10^13) para sacar codigos de seguimiento
		num = int64(source.Intn(10000000000000))

		//Manejar acceso concurrente al mapa para evitar errores de
		//sobre escritura o inconsistencias
		s.MuMapaSeg.Lock()
		_, err := s.MapaSeguimiento[num]
		s.MuMapaSeg.Unlock()
		if !err {
			s.MuMapaSeg.Lock()
			s.MapaSeguimiento[num] = "En bodega"
			s.MuMapaSeg.Unlock()
			break
		}
	}
	s.OrdenesP = append(s.OrdenesP, ordenP)
	return &Codigo{Cod: num}, err
}

//EnviarOrdenRetail para manejar las ordenes desde retail hasta servidor
func (s *Server) EnviarOrdenRetail(ctx context.Context, ordenR *OrdenRetail) (*Empty, error) {
	s.OrdenesR = append(s.OrdenesR, ordenR)
	var err error
	dummy := &Empty{}
	return dummy, err
}

// mustEmbedUnimplementedCourierServiceServer solo se añadio por compatibilidad
// y evitar warnings al compilar
func (s *Server) mustEmbedUnimplementedCourierServiceServer() {}
