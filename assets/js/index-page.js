function handleImage(img) {
  // If image is already loaded (from cache or fast connection)
  if (img.complete && img.naturalHeight !== 0) {    
    img.classList.remove("animate-pulse");
  } else {
    // Wait for it to load
    img.addEventListener("load", () => {
      console.log("Image loaded:", img.src);
      img.classList.remove("animate-pulse");
    });
  }
}

document.body.addEventListener("htmx:afterSwap", () => {
  const images = document.querySelectorAll("img");
  images.forEach(handleImage);
});


document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search)
  const q = params.get("q")

  if(q) {
    window.htmx.ajax('GET', `/search?q=${encodeURIComponent(q)}`, {target: '#posts-gallery'})
  } else {
    window.htmx.ajax('GET', '/latest-posts-with-img', {target: '#posts-gallery'})
  }

})