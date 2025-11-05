import handleImage from "./handle-image.js";
import loadPosts from "./loadPosts.js";
import "./keywords.js";

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
            }
        );
    } else {
        window.htmx.ajax("GET", "/posts-with-img/" + pageState.page, {
            target: "#posts-gallery",
        });
    }

    var select = new SlimSelect({
        select: "#keywords",
    });

    document
        .querySelector("#keywords")
        .closest("form")
        .addEventListener("reset", (e) => {
            // Updating SlimSelect from actual native select element value
            select.setSelected(
                Array.from(e.target.elements.select.selectedOptions).map(
                    (option) => option.value
                ),
                false
            );
        });
});

const loadMorePostsBtn = document.querySelector("#load-more-posts-btn");
loadMorePostsBtn.addEventListener("click", () => {
    loadPosts(pageState);
});
