export interface Blog {
    id?: number;
    title: string;
    slug?: string;
    content: string;
    image?: string;
    readingTime?: number;
    createdAt?: string;
};

export interface BlogResponse {
    data: Blog[];
    totalPages: number;
    currentPage: number;
}
