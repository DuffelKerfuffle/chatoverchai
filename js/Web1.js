

function display(title) {
  let children = document.getElementById(title).children;
  let displaySetting = "none";

  if (children[1].style.display === "none" ) {
    displaySetting = "block";
  }else if (children[1].style.display === "block"){
    displaySetting = "none";
  }
  hide(children, displaySetting);

  /*if (displayed != null && displayed != children) { 
    hide(displayed, "none")
  }
  
  let displaySetting = "none";
  
  if (children[1].style.display == "none" ) {
    displaySetting = "block";
  }

  if(displayed === null){
    displaySetting = "block";
  }

  hide(children, displaySetting);
  
  displayed = children;

  if (children[1].style.display == "none" ) {
    displayed = null;
  }*/
  
}

function load(){
  x = 1;
  //hide2(document.getElementById(0).children, "block");
  for(i = 0; i < 4; i++) {
    document.getElementById(0).children[i].style.display = "block";
  }
  while(document.getElementById(x) != null){
    //hide2(document.getElementById(x).children[0], "none");
    for(i = 0; i < 4; i++) {
      document.getElementById(x).children[i].style.display = "none";
    }
    x = x + 1;
  }
}

function hide(children, setting) {
  children[1].style.display = setting;
  children[2].style.display = setting;
}

function next(i){
  /*hide2(document.getElementById(i).children, "none");

  if(document.getElementById(i+1) != null){
    hide2(document.getElementById(i+1).children, "block");
  }else{
    hide2(document.getElementById(0).children, "block");
  }
  if(document.getElementById(i+2) != null){
    hide2(document.getElementById(i+2).children, "none");
  }*/

  x = 0;
  while(document.getElementById(x) != null){
    hide2(document.getElementById(x).children, "none");
    x = x + 1;
  }
  if(document.getElementById(i+1) != null){
    hide2(document.getElementById(i+1).children, "block");
  }else{
    hide2(document.getElementById(0).children, "block");
  }
}

function previous(i){
  /*hide2(document.getElementById(i).children, "none");
  
  if(document.getElementById(i-1) != null){
    hide2(document.getElementById(i-1).children, "block");
  }
  if(document.getElementById(i-2) != null){
    hide2(document.getElementById(i-2).children, "none");
  }*/
  x = 0;
  while(document.getElementById(x) != null){
    hide2(document.getElementById(x).children, "none");
    x = x + 1;
  }
  if(document.getElementById(i-1) != null){
    hide2(document.getElementById(i-1).children, "block");
  }else{
    hide2(document.getElementById(x-1).children, "block");
  }
}

function hide2(children, setting) {
  for(i = 0; i < 4; i++) {
    children[i].style.display = setting;
  }
}

function modify(){
  i = 0;
  while(document.getElementById("img" + i) != null){
    let str = document.getElementById("img" + i).src;
    str = str.split("/preview")[0];
    str = str.replace(".png","");
    document.getElementById("img"+i).src = str.split("/images/")[1];
    i++;
  }
}