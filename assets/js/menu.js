const menuToggler = document.querySelector("#menu-toggler")
const menu = document.querySelector("#menu")
menuToggler.addEventListener("click", () => {
    if (menu.classList.contains("translate-x-full")) {
        menu.classList.remove("translate-x-full")
    } else {
        menu.classList.add("translate-x-full")
    }
})