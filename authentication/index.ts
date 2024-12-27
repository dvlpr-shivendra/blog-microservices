interface User {
    id: string
    name: string
    email: string
    profilePicture: string
    createdAt: string
    updatedAt: string
}

const user: User = {
    id: "1",
    name: "Jake Roper",
    email: "jake@mailinator.com",
    profilePicture: "",
    createdAt: "",
    updatedAt: "",
}

console.log(user);