/**
 * Loads posts dynamically based on scroll position and query parameters
 * Uses HTMX to fetch and swap content inside the DOM
 *
 *@param {{page: number, isLoading: boolean}} pageState
 *
 */
function loadPosts(state) {
  const params = new URLSearchParams(window.location.search);
  const q = params.get("q");

  if (state.isLoading) return;
  state.isLoading = true;
  document.querySelector("#loader").classList.remove("hidden");
  const url = q
    ? `/search-posts-with-img?q=${encodeURIComponent(q)}&page=${state.page}`
    : `/posts-with-img/${state.page}`;

  fetch(url)
    .then((response) => {
      if (!response.ok) throw new Error("Network response was not ok");
      return response.text(); // Get raw HTML
    })
    .then((html) => {
      const container = document.querySelector("#posts-gallery");
      const fragment = document.createRange().createContextualFragment(html);
      container.appendChild(fragment); // Inject HTML partial
      // Manually trigger htmx:afterSwap
      const event = new CustomEvent("htmx:afterSwap", {
        bubbles: true,
        detail: {
          target: container,
          xhr: null, // optional, if you want to simulate the request
          swapSpec: { swapStyle: "beforeend" }, // optional
        },
      });
      container.dispatchEvent(event);

      state.page++;
      state.isLoading = false;
      document.querySelector("#loader").classList.add("hidden");
    })
    .catch((error) => {
      console.error("Error loading posts:", error);
      state.isLoading = false;
    });
}
export default loadPosts;
