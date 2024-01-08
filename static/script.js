function choosePhoto() {
    const fileInput = document.getElementById("photo");
    fileInput.click();
}
var toUploadFile;
function previewImage(event) {
    const input = event.target;
    const preview = document.getElementById("imagePreview");

    if (input.files && input.files[0]) {
        const reader = new FileReader();

        reader.onload = function (e) {
            preview.src = e.target.result;
            preview.style.display = "block";
        };

        reader.readAsDataURL(input.files[0]);
        toUploadFile = input.files[0];
    }
}

function uploadImage() {
    const form = document.getElementById("uploadForm");
    const formData = new FormData();
    const kValue = parseInt(document.getElementById("kValue").value);

    if (![2, 3, 4].includes(kValue)) {
        alert("Invalid value for 'k'. Please enter 2, 3, or 4.");
        return;
    }
    formData.append("image", toUploadFile);

    fetch(`http://localhost:1234/upload?k=${kValue}`, {
        method: "POST",
        body: formData,
    })
        .then((response) => response.json())
        .then((data) => displayImages(data.imagePaths))
        .catch((error) => console.error("Error:", error));
}


function displayImages(imagePaths) {
    if (!imagePaths) {
        console.error("Invalid response:", imagePaths);
        return;
    }

    const resultContainer = document.getElementById("result");
    resultContainer.innerHTML = ""; // Clear existing content

    // Ensure imagePaths is an array
    const pathsArray = Array.isArray(imagePaths) ? imagePaths : [imagePaths];

    // Create a container for the entire set of images
    const setContainer = document.createElement("div");
    setContainer.classList.add("centered-image-container");

    pathsArray.forEach((imagePath, index) => {
        const fullPath = "../" + imagePath;
        const container = document.createElement("div");
        container.classList.add("image-container");

        const image = document.createElement("img");
        image.src = fullPath;
        image.alt = `Image ${index + 1}`;

        // const closeButton = document.createElement("button");
        // closeButton.textContent = "Close";
        // closeButton.addEventListener("click", () => {
        //     container.remove();
        // });
        // container.appendChild(closeButton);

        container.appendChild(image);

        setContainer.appendChild(container);

        // Make the image draggable
        makeDraggable(container);
    });

    // Append the entire set container to the result container
    resultContainer.appendChild(setContainer);
    window.scrollTo({ top: 0, behavior: "smooth" });
}

function makeDraggable(element) {
    let offsetX, offsetY, isDragging = false;

    element.addEventListener("mousedown", (e) => {
        isDragging = true;
        offsetX = e.clientX - element.getBoundingClientRect().left;
        offsetY = e.clientY - element.getBoundingClientRect().top;
    });

    document.addEventListener("mousemove", (e) => {
        if (isDragging) {
            if (e.shiftKey) {
                // If Shift key is held down, snap to the top-left corner of the colliding image
                const collidingElement = getCollidingElement(element, e.clientX, e.clientY);
                if (collidingElement) {
                    const collidingRect = collidingElement.getBoundingClientRect();
                    element.style.left = collidingRect.left - element.clientLeft + "px";
                    element.style.top = collidingRect.top - element.clientTop + "px";
                    return;
                }
            }

            // Otherwise, move the element based on mouse movements
            element.style.left = e.clientX - offsetX + "px";
            element.style.top = e.clientY - offsetY + "px";
        }
    });

    document.addEventListener("mouseup", () => {
        isDragging = false;
    });
}

function getCollidingElement(element, mouseX, mouseY) {
    // Find the colliding element based on mouse position
    const elements = document.elementsFromPoint(mouseX, mouseY);
    return elements.find(el => el !== element && el.classList.contains("image-container"));
}
