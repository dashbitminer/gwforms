// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
  'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      validateFormat()
      if (!form.checkValidity()) {
        event.preventDefault()
        event.stopPropagation()
      }
    }, false)
  })
})()

  function activaselect1(){
    var select = document.getElementById("p47");
    var hiddenDiv = document.getElementById("afirmativo");
    var hiddenDiv2 = document.getElementById("negativo");
    
    // inputs afirmativos

    var input2 = document.getElementById("p59");
    var input3 = document.getElementById("p60");
    var input4 = document.getElementById("p61");
    var input5 = document.getElementById("p62");
    var input6 = document.getElementById("p63");
    var input7 = document.getElementById("p64");
    var input8 = document.getElementById("p65");
    var input9 = document.getElementById("p66");
    var input10 = document.getElementById("p67");
    var input11 = document.getElementById("p68");
    var input12 = document.getElementById("p69");
    var input13 = document.getElementById("p70");
    var input14 = document.getElementById("p71");
    var input15 = document.getElementById("p72");
    var input16 = document.getElementById("p73");
    var input17 = document.getElementById("fn1");
    var input18 = document.getElementById("fn2");
    var input19 = document.getElementById("fn3");
    var input20 = document.getElementById("fn4");
    var input21 = document.getElementById("p58");

    var input23 = document.getElementById("p57");

    // inputs negativos
    var input24 = document.getElementById("p50");

    if(select.value==="1"){
      hiddenDiv.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none");
   
      input2.required=true;
      input3.required=true;
      input4.required=true;
      input5.required=true;
      input6.required=true;
      input7.required=true;
      input8.required=true;
      input9.required=true;
      input10.required=true;
      input11.required=true;
      input12.required=true;
      input13.required=true;
      input14.required=true;
      input15.required=true;
      input16.required=true;
      input17.required=true;
      input18.required=true;
      input19.required=true;
      input20.required=true;
      input21.required=true;
     
      input23.required=true;
      input24.required=false;      
    }
    else{
      hiddenDiv.classList.add("d-none");
      hiddenDiv2.classList.remove("d-none");
     
      input2.required=false;
      input3.required=false;
      input4.required=false;
      input5.required=false;
      input6.required=false;
      input7.required=false;
      input8.required=false;
      input9.required=false;
      input10.required=false;
      input11.required=false;
      input12.required=false;
      input13.required=false;
      input14.required=false;
      input15.required=false;
      input16.required=false;
      input17.required=false;
      input18.required=false;
      input19.required=false;
      input20.required=false;
      input21.required=false;
    
      input23.required=false;
      input24.required=true;
    }
  }

  function activaselect2(){
    var select = document.getElementById("p58");
    var hiddenDiv = document.getElementById("afirmativo2");
    // inputs afirmativos
    var input = document.getElementById("nota");
    var input2 = document.getElementById("xq")
  
    if(select.value==="1"){
      hiddenDiv.classList.remove("d-none");
      input.required=true;
      input2.required=true;
    }
    else{
      hiddenDiv.classList.add("d-none");
      input.required=false;
      input2.required=false;
    }
  }

  function activaselect3(){
    var select = document.getElementById("p57");
    var hiddenDiv = document.getElementById("afirmativo3");
    // inputs firmativos
    var input = document.getElementById("email");
  
    if(select.value==="1"){
      hiddenDiv.classList.remove("d-none");
      input.required=true;      
    }
    else{
      hiddenDiv.classList.add("d-none");
      input.required=false;
    }
  }

  function activaselect4(){
    var select = document.getElementById("p50");
    var hiddenDiv = document.getElementById("afirmativo4");
    var hiddenDiv2 = document.getElementById("negativo2");

    // inputs afirmativos
    var input = document.getElementById("necesidad");
  
    // inputs negativos
   

    if(select.value==="1"){
      hiddenDiv.classList.remove("d-none");
      hiddenDiv2.classList.add("d-none"); 
      input.required=true;     

    }
    else{
      hiddenDiv.classList.add("d-none");
      hiddenDiv2.classList.remove("d-none");
      input.required=false;

    }
  }

  function activaselect5(){  
    var hiddenDiv = document.getElementById("otrarazon");
    var input = document.getElementById("otrazon")
  
    var checkboxes = document.querySelectorAll('input[name="p52"]');
for (var i = 0; i < checkboxes.length; i++) {
  if(checkboxes[i].value==="88" && checkboxes[i].checked){
    hiddenDiv.classList.remove("d-none");
    input.required=true;
  }
  else{
    hiddenDiv.classList.add("d-none");
    input.required=false;
  }
}

  }

  function activaselect6(){
    var select = document.getElementById("p46");
    var hiddenDiv = document.getElementById("otrorol");
    var input = document.getElementById("otrol")
  
    if(select.value==="80"){
      hiddenDiv.classList.remove("d-none");
      input.required=true;
    }
    else{
      hiddenDiv.classList.add("d-none");
      input.required=false;
    }
  }

  function activasugerencia1(){
    var select = document.getElementById("p59");
    var hiddenDiv = document.getElementById("ncf1");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia2(){
    var select = document.getElementById("p60");
    var hiddenDiv = document.getElementById("ncf2");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia3(){
    var select = document.getElementById("p61");
    var hiddenDiv = document.getElementById("ncf3");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia4(){
    var select = document.getElementById("p62");
    var hiddenDiv = document.getElementById("ncf4");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia5(){
    var select = document.getElementById("p63");
    var hiddenDiv = document.getElementById("ncf5");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia6(){
    var select = document.getElementById("p64");
    var hiddenDiv = document.getElementById("ncf6");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia7(){
    var select = document.getElementById("p65");
    var hiddenDiv = document.getElementById("ncf7");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia8(){
    var select = document.getElementById("p66");
    var hiddenDiv = document.getElementById("ncf8");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia9(){
    var select = document.getElementById("p67");
    var hiddenDiv = document.getElementById("ncf9");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia10(){
    var select = document.getElementById("p68");
    var hiddenDiv = document.getElementById("ncf10");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia11(){
    var select = document.getElementById("p69");
    var hiddenDiv = document.getElementById("ncf11");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia12(){
    var select = document.getElementById("p70");
    var hiddenDiv = document.getElementById("ncf12");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia13(){
    var select = document.getElementById("p71");
    var hiddenDiv = document.getElementById("ncf13");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia14(){
    var select = document.getElementById("p72");
    var hiddenDiv = document.getElementById("ncf14");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function activasugerencia15(){
    var select = document.getElementById("p73");
    var hiddenDiv = document.getElementById("ncf15");
  
    if(select.value==="94"){
      hiddenDiv.classList.add("d-none");
      
    }
    else{
      hiddenDiv.classList.remove("d-none");
      
    }
  }

  function validarEmail() {
    const inputEmail = document.getElementById("email").value;
    
    // Expresión regular para validar el formato de correo electrónico
    const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
    
    // Lista de dominios permitidos
    const dominiosPermitidos = ["gmail.com", "hotmail.com", "outlook.com", "yahoo.com", "crisalidainternational.org", "glasswing.org"];
    
    const dominio = inputEmail.split('@')[1];
    const errorEmail = document.getElementById("emailHelp");
    
    if (!emailPattern.test(inputEmail)) {
        errorEmail.textContent = "Formato de correo electrónico inválido";
        errorEmail.style.color = "red"; // Cambiar el color del mensaje a rojo
    } else if (dominiosPermitidos.indexOf(dominio) === -1) {
        errorEmail.textContent = "Dominio no permitido";
        errorEmail.style.color = "red"; // Cambiar el color del mensaje a rojo
    } else {
        errorEmail.textContent = ""; // Reiniciar mensaje de error
    }
}