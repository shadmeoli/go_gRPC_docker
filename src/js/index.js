function openModal() {
  const modal = document.getElementById("newNoteModal");
  if (modal) {
    modal.classList.remove("hidden");
  }
}

function closeModal() {
  const modal = document.getElementById("newNoteModal");
  if (modal) {
    modal.classList.add("hidden");
  }
}

function saveNote() {
  const noteText = document.getElementById("noteText").value;
  // Perform actions to save the new note, e.g., sending it to the server
  console.log("New note saved:", noteText);
  closeModal();
}

document.addEventListener("keydown", function (event) {
  if (event.altKey && event.key === "n") {
    // Alt + n is pressed, execute New note action
    event.preventDefault(); // Prevent the default browser behavior
    openModal();
  } else if (event.altKey && event.key === "s") {
    // Alt + s is pressed, execute Save action
    event.preventDefault(); // Prevent the default browser behavior
    saveNote();
  } else if (event.altKey && event.key === "x") {
    // Alt + x is pressed, execute Delete note action
    event.preventDefault(); // Prevent the default browser behavior
    console.log("Delete note action");
  }
});
