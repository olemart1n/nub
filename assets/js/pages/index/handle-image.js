function handleImage(img) {
  // If image is already loaded (from cache or fast connection)
  if (img.complete && img.naturalHeight !== 0) {
    img.parentElement.classList.remove("animate-pulse");
  } else {
    // Wait for it to load
    img.addEventListener("load", () => {
      img.parentElement.classList.remove("animate-pulse");
    });
  }
}

export default handleImage;
