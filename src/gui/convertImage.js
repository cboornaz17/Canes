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
}
function failed() {
  console.error("The provided file couldn't be loaded as an Image media");
}

document.getElementById('submit').onclick = function(e) {
  var canvas = document.getElementById('canvas');
  var img64 = getImg64(canvas);
  post("/images/", {data: 'img64'});
}
function getImg64(canvas) {
  return canvas.toDataURL('image/jpeg', 1.0)

}
  /**
  * sends a request to the specified url from a form. this will change the window location.
  * @param {string} path the path to send the post request to
  * @param {object} params the paramiters to add to the url
  * @param {string} [method=post] the method to use on the form
  */

 function post(path, params, method) {
     method = method || "post"; // Set method to post by default if not specified.

     // The rest of this code assumes you are not using a library.
     // It can be made less wordy if you use one.
     var form = document.createElement("form");
     form.setAttribute("method", method);
     form.setAttribute("action", path);

     for(var key in params) {
         if(params.hasOwnProperty(key)) {
             var hiddenField = document.createElement("input");
             hiddenField.setAttribute("type", "hidden");
             hiddenField.setAttribute("name", key);
             hiddenField.setAttribute("value", params[key]);

             form.appendChild(hiddenField);
         }
     }

     document.body.appendChild(form);
     form.submit();
 }
