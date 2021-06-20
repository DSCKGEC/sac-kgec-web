const Slidingnav = () => {

    const burger = document.querySelector(".burger");
    const nav = document.querySelector(".navlinks");
    const navlinks = document.querySelector(".navlinks li");
    // Toggle nav
    burger.addEventListener("click", () => {
        nav.classList.toggle("navactive");
    });
    // animate links
    // navlinks.forEach((link, index) => { console.log(index); });

}

Slidingnav();