// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'
  
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')
  
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
      form.addEventListener('submit', event => {
      document.getElementById("boton").Disabled = true; 
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
          document.getElementById("boton").Disabled = false;
        }
      }, false)
    })
  })()

function actualizarMoneda(){
  const pais = document.getElementById("pais").value;
  const textoMoneda = document.getElementById("textoMoneda");

  const textosPorPais = {
    7: "Como incentivo y agradecimiento a su tiempo en el llenado de esta encuesta, usted estará participando en la rifa de cinco tarjetas de regalo de $50.00 en Walmart. La rifa se realizará entre el 16 y 17 de diciembre de 2024. Únicamente cinco personas serán ganadoras de todas las participantes en la encuesta.",
    4: "Como incentivo y agradecimiento a su tiempo en el llenado de esta encuesta, usted estará participando en la rifa de cinco tarjetas de regalo de Q385.00 en Walmart. La rifa se realizará entre el 16 y 17 de diciembre de 2024. Únicamente cinco personas serán ganadoras de todas las participantes en la encuesta.",
    3: "Como incentivo y agradecimiento a su tiempo en el llenado de esta encuesta, usted estará participando en la rifa de cinco tarjetas de regalo de L 1,250.00 en Walmart. La rifa se realizará entre el 16 y 17 de diciembre de 2024. Únicamente cinco personas serán ganadoras de todas las participantes en la encuesta."
  };

  textoMoneda.textContent = textosPorPais[pais];

}

function activaselect(){
var select = document.getElementById("participar");
var hiddenDiv = document.getElementById("ocultar"); 
var input1=document.getElementById("trabajo");
var input2=document.getElementById("edad");
var input3=document.getElementById("aniol");
var input4=document.getElementById("sexo");
var input5=document.getElementById("rol");
var input6=document.getElementById("afecta1");
var input7=document.getElementById("afecta2");
var input8=document.getElementById("afecta3");
var input9=document.getElementById("actitud");
var input10=document.getElementById("estrategia48");
var input11=document.getElementById("practicat");
var input14=document.getElementById("opinionAtend");
var input32=document.getElementById("porcentaje");
var input33=document.getElementById("82");
var input34=document.getElementById("practicat1242");
var input36=document.getElementById("afirmaciones1");
var input37=document.getElementById("afirmaciones2");
var input38=document.getElementById("afirmaciones3");
var input40=document.getElementById("utilidad242");
var input42=document.getElementById("resolucion242");
var input43=document.getElementById("conocia242");
var input44=document.getElementById("consciente242");
var input45=document.getElementById("beneficio");
var input46=document.getElementById("rel1");
var input47=document.getElementById("rel2");
var input48=document.getElementById("rel3");
var input49=document.getElementById("rel4");
var input50=document.getElementById("rel12");
var input51=document.getElementById("rel22");
var input52=document.getElementById("rel32");
var input53=document.getElementById("rel42");




if(select.value==="227"){ 
hiddenDiv.classList.remove("d-none");
input1.required = true
input2.required = true
input3.required = true
input4.required = true
input5.required = true
input6.required = true
input7.required = true
input8.required = true
input9.required = true
input10.required = true
input11.required = true
input14.required = true
input32.required = true
input33.required = true
input34.required = true
input36.required = true
input37.required = true
input38.required = true
input40.required = true
input42.required = true
input43.required = true
input44.required = true
input45.required = true
input46.required = true
input47.required = true
input48.required = true
input49.required = true
input50.required = true
input51.required = true
input52.required = true
input53.required = true

}
else{
  hiddenDiv.classList.add("d-none");
input1.required = false
input2.required = false
input3.required = false
input4.required = false
input5.required = false
input6.required = false
input7.required = false
input8.required = false
input9.required = false
input10.required = false
input11.required = false
input14.required = false
input32.required = false
input33.required = false
input34.required = false
input36.required = false
input37.required = false
input38.required = false
input40.required = false
input42.required = false
input43.required = false
input44.required = false
input45.required = false
input46.required = false
input47.required = false
input48.required = false
input49.required = false
input50.required = false
input51.required = false
input52.required = false
input53.required = false

}
}

function otro(){
  var select=document.getElementById('estrategia48');
  var select2=document.getElementById('estrategia49');
  var select3=document.getElementById('estrategia50');
  var hiddenDiv4=document.getElementById('epracticadiv');
  var input=document.getElementById('aplicadoP');
  if (select.checked){
    hiddenDiv4.classList.remove("d-none");
    input.required = true;
  }
  
  if (select2.checked){
    hiddenDiv4.classList.remove("d-none");
    input.required = true;
  }
  
  if (select3.checked){
    hiddenDiv4.classList.add("d-none");
    input.required = false;
  }
}

function otro4(){
  var select3=document.getElementById('apoyo241');
  var hiddenDiv5=document.getElementById('motivodiv');
  var input=document.getElementById('motivo');
  if (select3.checked){
    hiddenDiv5.classList.remove("d-none");
    input.required = true;
  }
  else{
    hiddenDiv5.classList.add("d-none");
    input.required = false
    input.value = "";
  }
}

function otro5(){
  var select4=document.getElementById('utilidad243');
  var hiddenDiv5=document.getElementById('ejemplosl');
  var input=document.getElementById('ejemplos');
  if (select4.checked){
    hiddenDiv5.classList.add("d-none");
    input.required = false;
  }
  else{
    hiddenDiv5.classList.remove("d-none");
    input.required = true;
  }
}

function otro6(){
  var select=document.getElementById('practicat1242');
  var hiddenDiv4=document.getElementById('831div');
  var input=document.getElementById('motivo2');
  if (select.checked){
    hiddenDiv4.classList.remove("d-none");
    input.required = true;
  }
  else{
    hiddenDiv4.classList.add("d-none");
    input.required = false;
  }

}

function otro3(checkbox){
  var hiddenDiv1 = document.getElementById("apoyodiv");
  var input = document.getElementById("apoyo238");
  var errorMensaje = document.getElementById("errorMensaje1");
  var hiddenDiv2 = document.getElementById("otraEstrategiadiv");
  var input1 = document.getElementById("otraEstrategia");

  // Comprobar si el checkbox con valor 4 está seleccionado
  var isChecked = document.querySelector('input[name="epractica"][value="56"]').checked;
  var isChecked1 = document.querySelector('input[name="epractica"][value="57"]').checked;

  // Comprobar si al menos un checkbox está seleccionado
  var anyChecked = document.querySelectorAll('input[name="epractica"]:checked').length > 0;

  if (isChecked) {
    hiddenDiv1.classList.remove("d-none");
    input.required = true;
  } else {
    hiddenDiv1.classList.add("d-none");
    input.required = false;
    input.value = ""; // Resetea el valor del input
  }
  
  if (isChecked1) {
    hiddenDiv2.classList.remove("d-none");
    input1.required = true;
  } else {
    hiddenDiv2.classList.add("d-none");
    input1.required = false;
    input1.value = ""; // Resetea el valor del input
  }

// Mostrar u ocultar el mensaje de error
  if (anyChecked) {
      errorMensaje.style.display = "none"; // Ocultar mensaje de error
  } else {
      errorMensaje.style.display = "block"; // Mostrar mensaje de error
  }
}

function afirmacionesP(checkbox){
  var errorMensaje = document.getElementById("errorMensaje2");

  // Comprobar si al menos un checkbox está seleccionado
  var anyChecked = document.querySelectorAll('input[name="op101"]:checked').length > 0;

// Mostrar u ocultar el mensaje de error
  if (anyChecked) {
      errorMensaje.style.display = "none"; // Ocultar mensaje de error
  } else {
      errorMensaje.style.display = "block"; // Mostrar mensaje de error
  }
}

function perfil7(){
  var select=document.getElementById('rol');
  var hiddenDiv1=document.getElementById('perfilDocente');
  var hiddenDiv2=document.getElementById('perfilP');
  var input12=document.getElementById("personasAt1");
  var input13=document.getElementById("personasAt2");

  if (select.value==="229"){
    hiddenDiv1.classList.remove("d-none");
    hiddenDiv2.classList.add("d-none");
    input12.required = false;
    input13.required = true;
 }
  else{
    hiddenDiv1.classList.add("d-none");
    hiddenDiv2.classList.remove("d-none");
    input12.required = true;
    input13.required = false;
  }
}

document.addEventListener("DOMContentLoaded", function() {
  const images = document.querySelectorAll(".img-max");

  images.forEach(function(image) {
    image.addEventListener("click", function() {
      const radio = image.closest(".form-check").querySelector(".form-check-input");
      radio.checked = true;
    });
  });
});

document.addEventListener('DOMContentLoaded', function () {
  const bar = document.getElementById('percentageBar');
  const points = bar.querySelectorAll('.percentage-point');
  let selectedValue = null; // Variable para almacenar el valor seleccionado

  points.forEach((point, index) => {
      // Resaltar puntos al pasar el mouse
      point.addEventListener('mouseover', function () {
          points.forEach((p, i) => {
              if (i <= index) {
                  p.classList.add('highlight');
              } else {
                  p.classList.remove('highlight');
              }
          });
      });

      // Seleccionar un punto al hacer clic
      point.addEventListener('click', function () {
          points.forEach(p => p.classList.remove('selected'));
          for (let i = 0; i <= index; i++) {
              points[i].classList.add('selected');
          }
          selectedValue = point.getAttribute('data-value'); // Almacenar el valor seleccionado
          document.getElementById("porcentaje").value = selectedValue
      });
  });

  // Limpiar resaltado cuando el mouse sale de la barra
  bar.addEventListener('mouseleave', function () {
      points.forEach(point => point.classList.remove('highlight'));
  });
});

function frecuencia() {
  var select1 = document.getElementById('rol');
  var selectedRadio = document.querySelector('input[name="opinionAtend"]:checked');
  var hiddenDiv1 = document.getElementById('personalDocente');
  var hiddenDiv2 = document.getElementById('personalSPO');
  var hiddenDiv5 = document.getElementById('porcentajediv');
  var input = document.getElementById('porcentaje');
  var input12=document.getElementById("frecuencia1");
  var input13=document.getElementById("frecuencia2");
  var input14=document.getElementById("frecuencia3");
  var input15=document.getElementById("frecuencia4");
  var input16=document.getElementById("frecuencia5");
  var input17=document.getElementById("frecuencia6");
  var input18=document.getElementById("frecuencia7");
  var input19=document.getElementById("frecuencia8");
  var input20=document.getElementById("frecuencia9");
  var input21=document.getElementById("frecuencia11");
  var input22=document.getElementById("frecuencia12");
  var input23=document.getElementById("frecuencia13");
  var input24=document.getElementById("frecuencia14");
  var input25=document.getElementById("frecuencia15");
  var input26=document.getElementById("frecuencia16");
  var input27=document.getElementById("frecuencia17");
  var input28=document.getElementById("frecuencia18");
  // Agrega más inputs según sea necesario

  // Validar si hay un radio seleccionado
  var opinionValue = selectedRadio ? selectedRadio.value : "";

  // Condiciones
  if (select1.value === "229" && opinionValue === "245") {
      hiddenDiv1.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv5.classList.remove("d-none");
      input.required = true;
      input12.required = false;
      input13.required = false;
      input14.required = false;
      input15.required = false;
      input16.required = false;
      input17.required = false;
      input18.required = false;
      input19.required = false;
      input20.required = false;
      input21.required = true;
      input22.required = true;
      input23.required = true;
      input24.required = true;
      input25.required = true;
      input26.required = true;
      input27.required = true;
      input28.required = true;
      // Configura más inputs según sea necesario
  } else if (select1.value !== "229" && opinionValue === "245") {
      hiddenDiv1.classList.add("d-none");
      hiddenDiv2.classList.remove("d-none");
      hiddenDiv5.classList.remove("d-none");
      input.required = true;
      input12.required = true;
      input13.required = true;
      input14.required = true;
      input15.required = true;
      input16.required = true;
      input17.required = true;
      input18.required = true;
      input19.required = true;
      input20.required = true;
      input21.required = false;
      input22.required = false;
      input23.required = false;
      input24.required = false;
      input25.required = false;
      input26.required = false;
      input27.required = false;
      input28.required = false;
      // Configura más inputs según sea necesario
  } else if(select1.value === "229" && opinionValue === "244"){
      hiddenDiv1.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv5.classList.remove("d-none");
      input.required = true;
      input12.required = false;
      input13.required = false;
      input14.required = false;
      input15.required = false;
      input16.required = false;
      input17.required = false;
      input18.required = false;
      input19.required = false;
      input20.required = false;
      input21.required = true;
      input22.required = true;
      input23.required = true;
      input24.required = true;
      input25.required = true;
      input26.required = true;
      input27.required = true;
      input28.required = true;
    }else if (select1.value !== "229" && opinionValue === "244") {
      hiddenDiv1.classList.add("d-none");
      hiddenDiv2.classList.remove("d-none");
      hiddenDiv5.classList.remove("d-none");
      input.required = true;
      input12.required = true;
      input13.required = true;
      input14.required = true;
      input15.required = true;
      input16.required = true;
      input17.required = true;
      input18.required = true;
      input19.required = true;
      input20.required = true;
      input21.required = false;
      input22.required = false;
      input23.required = false;
      input24.required = false;
      input25.required = false;
      input26.required = false;
      input27.required = false;
      input28.required = false;
      // Configura más inputs según sea necesario
    } else if(opinionValue === "246") {
      hiddenDiv1.classList.add("d-none");
      hiddenDiv2.classList.add("d-none");
      hiddenDiv5.classList.add("d-none");
      input.required = false;
      input12.required = false;
      input13.required = false;
      input14.required = false;
      input15.required = false;
      input16.required = false;
      input17.required = false;
      input18.required = false;
      input19.required = false;
      input20.required = false;
      input21.required = false;
      input22.required = false;
      input23.required = false;
      input24.required = false;
      input25.required = false;
      input26.required = false;
      input27.required = false;
      input28.required = false;
      // Configura más inputs según sea necesario
  } 
}

document.getElementById('formualrioRegistro').addEventListener('submit', function(event) {
  var checkboxes = document.querySelectorAll('input[name="epractica"]');
  var alMenosUnoMarcado = false;
  var select = document.getElementById('estrategia50');
  var select1 = document.getElementById("participar");
  var errorMensaje = document.getElementById('errorMensaje1');

  if (select.checked){
    errorMensaje.style.display = 'none';
    return;
  } else if (select1.value==="228"){
    errorMensaje.style.display = 'none';
    return;
  } else {
    for (var i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i].checked) {
            alMenosUnoMarcado = true;
            break;
        }
    }

    if (!alMenosUnoMarcado) {
      event.preventDefault(); // Evitar el envío del formulario
      errorMensaje.style.display = 'block'; // Mostrar mensaje de error
      errorMensaje.focus();
      document.getElementById("boton").Disabled = false; // Hacer foco en el mensaje de error
    } else {
      document.getElementById('errorMensaje1').style.display = 'none'; // Ocultar mensaje de error si está marcado
    }
  }
});

document.getElementById('formualrioRegistro').addEventListener('submit', function(event) {
  var checkboxes = document.querySelectorAll('input[name="op101"]');
  var alMenosUnoMarcado = false;
  var select1 = document.getElementById("participar");
  var errorMensaje2 = document.getElementById('errorMensaje2');

  if (select1.value==="228"){
    errorMensaje2.style.display = 'none';
    return;
  } else {
    for (var i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i].checked) {
            alMenosUnoMarcado = true;
            break;
        }
    }

    if (!alMenosUnoMarcado) {
      event.preventDefault(); // Evitar el envío del formulario    
      errorMensaje2.style.display = 'block'; // Mostrar mensaje de error
      errorMensaje2.focus(); // Hacer foco en el mensaje de error
      document.getElementById("boton").Disabled = false; 
    } else {
      document.getElementById('errorMensaje2').style.display = 'none'; // Ocultar mensaje de error si está marcado
    }
  }
});

