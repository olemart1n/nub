// assets/js/pages/index/keywords.js

const select = document.getElementById("keywords"); // <select> element
const selectedKeywordsDiv = document.getElementById("selected-keywords"); // <div> element

// Function to update the URL with the selected keywords
function updateURL(keywords) {
    const url = new URL(window.location.href);
    url.searchParams.set("keywords", keywords);
    window.history.pushState({}, "", url);
}

// Function to fetch data from the server
async function fetchData(keywords) {
    try {
        const url = new URL(window.location.href);
        url.searchParams.set("keyword", keywords);
        const response = await fetch(url);
        const data = await response.json();
        console.log("Data fetched:", data);
        // Handle the fetched data as needed
    } catch (error) {
        console.error("Error fetching data:", error);
    }
}

// Event listener for the change event on the select element
select.addEventListener("change", function () {
    const selectedOption = select.selectedOptions[0].value;

    // Update the URL with the selected keywords
    updateURL(selectedOption);
});
