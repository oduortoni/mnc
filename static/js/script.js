let exploreBtn; // explore button

window.onload = function() {
    exploreBtn = document.querySelector("#explore");
    exploreBtn.addEventListener("click", explore);
}

// toggle navbar
const navbarToggle = document.querySelector('.navbar-toggle');
const navbarLinks = document.querySelector('.navbar-links');
navbarToggle.addEventListener('click', () => {
    navbarLinks.classList.toggle('active');
});
  
/*
* used to help the user explore the list of available rooms
*/
async function explore() {
    try {
        let rooms = await fetch("/explore")
            .then(result => result.json())
            .then(data => data);

        console.log("Rooms:", rooms);

        // let roomsHtml = JSON.stringify(rooms, null, 2);
        // console.log("RoomsTML: ", roomsHtml)

        let roomsListHtml = '<ul>';
        rooms.rooms.forEach(room => {
            roomsListHtml += `<li>Room: ${room.name}, Capacity: ${room.capacity}</li>`;
        });
        roomsListHtml += '</ul>';

        // Insert the roomsListHtml into the #exploration element
        document.querySelector("#exploration").innerHTML = roomsListHtml;
    } catch (error) {
        console.error("Error fetching rooms:");
        document.querySelector("#exploration").innerHTML = "Failed to load rooms data.";
    }
}
