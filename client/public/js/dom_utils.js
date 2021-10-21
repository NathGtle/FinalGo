
// Code repris de l'utilisateur stackoverflow : jfriend00 
// Depuis : https://stackoverflow.com/questions/9899372/pure-javascript-equivalent-of-jquerys-ready-how-to-call-a-function-when-t
function OnDocReady(fn) {
	if (document.readyState === "complete" || 
    	document.readyState === "interactive") {
        setTimeout(fn, 1);
    } else {
        document.addEventListener("DOMContentLoaded", fn);
    }
}

