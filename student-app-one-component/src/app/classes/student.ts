export interface book{
    id:string,
    name:string,
    totalStock:number,
}

export interface bookIssues{
    id:string,
    bookId:string,
    studentId:string,
    book:book,
    issueDate:string,
    returned:boolean
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
    bookIssues:bookIssues[]
}
