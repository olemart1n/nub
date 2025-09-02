import handleImage from "./handle-image.js";
import loadPosts from "./loadPosts.js";

const pageState = { page: 0, isLoading: false };

document.body.addEventListener("htmx:afterSwap", () => {
  const images = document.querySelectorAll("img");
  images.forEach(handleImage);
});

document.addEventListener("DOMContentLoaded", () => {
  const params = new URLSearchParams(window.location.search);
  const q = params.get("q");

  if (q) {
    window.htmx.ajax(
      "GET",
      `/search-posts-with-img?q=${encodeURIComponent(q)}`,
      {
        target: "#posts-gallery",
      },
    );
  } else {
    window.htmx.ajax("GET", "/posts-with-img/" + pageState.page, {
      target: "#posts-gallery",
    });
  }
});

const loadMorePostsBtn = document.querySelector("#load-more-posts-btn");
loadMorePostsBtn.addEventListener("click", () => {
  loadPosts(pageState);
});
