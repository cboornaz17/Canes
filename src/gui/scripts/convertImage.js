document.getElementById('newImage').onchange = function(e) {
  var img = new Image();
  img.onload = draw;
  img.onerror = failed;
  img.src = URL.createObjectURL(this.files[0]);
};
function draw() {
  var canvas = document.getElementById('canvas');
  canvas.width = this.width;
  canvas.height = this.height;
  var ctx = canvas.getContext('2d');
  ctx.drawImage(this, 0,0);
  var img64 = getImg64(canvas);
  console.log(img64);
}
function failed() {
  console.error("The provided file couldn't be loaded as an Image media");
}
function getImg64(canvas) {
  return canvas.toDataURL('image/jpeg', 1.0)
}
