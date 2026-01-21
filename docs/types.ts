// types.ts

interface User {
  id: number;
  name: string;
  email: string;
}

interface Product {
  id: number;
  name: string;
  price: number;
}

interface Order {
  id: number;
  userId: number;
  productId: number;
  quantity: number;
  total: number;
  status: string;
}

interface Filter {
  limit: number;
  offset: number;
}

interface FilterParams {
  limit?: number;
  offset?: number;
  order?: string[];
  sort?: string;
}

interface OrderParams {
  userId: number;
  productId: number;
  quantity: number;
  status: string;
}

interface Response<T> {
  data: T[];
  meta: {
    total: number;
    limit: number;
    offset: number;
  };
}

interface ErrorResponse {
  message: string;
  status: number;
}