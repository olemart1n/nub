function handleImage(img) {
  // If image is already loaded (from cache or fast connection)
  if (img.complete && img.naturalHeight !== 0) {
    console.log("Image already loaded:", img.src);
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