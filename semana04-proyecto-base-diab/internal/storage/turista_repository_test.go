// Turista
package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestTuristaGuardar_TableDriven(t *testing.T) {
	repo := NewTuristaMemoria()

	// Pre-condición: sembramos un negocio para poder probar "ID duplicado".
	turistaBase := models.Turista{
		ID: 1, Nombre: "Juan",
		Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
	}
	if err := repo.Guardar(turistaBase); err != nil {
		// t.Fatalf detiene el test inmediatamente. Si el setup falla,
		// no tiene sentido seguir corriendo el resto de los casos.
		t.Fatalf("setup falló: %v", err)
	}

	// La tabla de casos. Cada elemento es un escenario completo:
	// nombre del subtest, datos de entrada, error esperado.
	casos := []struct {
		nombre    string
		entrada   models.Turista
		esperaErr error
	}{
		{
			nombre: "caso feliz - Turista válido",
			entrada: models.Turista{
				ID: 100, Nombre: "Café Manabita",
				Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
			},
			esperaErr: nil,
		},
		{
			nombre: "nombre vacío falla",
			entrada: models.Turista{
				ID: 101, Nombre: "",
				Nacionalidad: "Española", IdiomaPreferido: "es",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "nacinalidad vacia",
			entrada: models.Turista{
				ID: 102, Nombre: "Jose",
				Nacionalidad: "", IdiomaPreferido: "en",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "Idioma Prefefiro no valido",
			entrada: models.Turista{
				ID: 102, Nombre: "Pepe",
				Nacionalidad: "Peruana", IdiomaPreferido: "ent",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.Turista{
				ID: 1, Nombre: "Otro nombre",
				Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	// Iteramos sobre los casos y corremos un subtest por cada uno.
	// t.Run permite que cada subtest se reporte por separado y que se
	// puedan correr individualmente con `go test -run`.
	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {
			err := repo.Guardar(c.entrada)

			// errors.Is es la forma idiomática de comparar errores
			// tipados. NUNCA uses err == c.esperaErr ni
			// err.Error() == "..." — son frágiles.
			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(%q): esperaba error=%v, obtuvo error=%v",
					c.entrada.Nombre, c.esperaErr, err)
			}
		})
	}
}

// TestBuscarPorID_NegocioExiste verifica el caso feliz de BuscarPorID.
//
// Este es un test SIMPLE de un solo caso. No necesita el patrón
// table-driven porque solo hay un comportamiento esperado a verificar.
//
// Los OTROS casos de BuscarPorID (ID negativo, ID inexistente) deberían
// ir en otro test, posiblemente table-driven, que VOS tenés que escribir.
func TestBuscarPorID_TuristaExiste(t *testing.T) {
	repo := NewTuristaMemoria()

	// Arrange: creamos y guardamos un negocio.
	esperado := models.Turista{
		ID: 42, Nombre: "Juan",
		Nacionalidad: "Ecuatoriana", IdiomaPreferido: "es",
	}
	if err := repo.Guardar(esperado); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	// Act: buscamos el negocio por su ID.
	obtenido, err := repo.BuscarPorID(42)

	// Assert: no debe haber error y debe coincidir con lo guardado.
	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if obtenido.ID != esperado.ID {
		t.Errorf("ID: esperaba %d, obtuvo %d", esperado.ID, obtenido.ID)
	}
	if obtenido.Nombre != esperado.Nombre {
		t.Errorf("Nombre: esperaba %q, obtuvo %q", esperado.Nombre, obtenido.Nombre)
	}
	if obtenido.Nacionalidad != esperado.Nacionalidad {
		t.Errorf("Nacionalidad: esperaba %q, obtuvo %q", esperado.Nacionalidad, obtenido.Nacionalidad)
	}
	if obtenido.IdiomaPreferido != esperado.IdiomaPreferido {
		t.Errorf("IdiomaPreferido: esperaba %q, obtuvo %q", esperado.IdiomaPreferido, obtenido.IdiomaPreferido)
	}
}
