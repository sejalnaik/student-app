export interface Book{
    id:string,
    name:string,
    totalStock:number,
}

export interface BookWithAvailable{
    id:string,
    name:string,
    totalStock:number,
    available:number
}

export interface BookIssues{
    id:string,
    bookId:string,
    studentId:string,
    book:Book,
    issueDate:string,
    returned:boolean,
    penalty:number
}

export interface Student {
    id:string,
    name:string,
    rollNo:number,
    age:number,
    dob:string,
    dobTime : string
    email:string,
    isMale:boolean
    phoneNumber:string
}
