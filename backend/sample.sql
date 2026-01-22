-- Cart
CONSTRAINT fk_cart_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
CONSTRAINT fk_cart_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE

-- Orders
CONSTRAINT fk_order_user FOREIGN KEY (user_id) REFERENCES users(id)

-- Order Details
CONSTRAINT fk_order_detail_order FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE
CONSTRAINT fk_order_detail_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE,

-- Order Returns
CONSTRAINT fk_order_return_order FOREIGN KEY(order_id) REFERENCES orders(id) ON DELETE CASCADE,
CONSTRAINT fk_order_return_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE

-- Order Return Details
CONSTRAINT fk_order_return_detail_order_return FOREIGN KEY(order_return_id) REFERENCES order_returns(id) ON DELETE CASCADE,
CONSTRAINT fk_order_return_detail_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE

-- Products
CONSTRAINT fk_product_category FOREIGN KEY(category_id) REFERENCES categories(id) ON DELETE CASCADE,
CONSTRAINT fk_product_supplier FOREIGN KEY(supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE

-- Purchase
CONSTRAINT fk_purchase_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE

-- Purchase Details
CONSTRAINT fk_purchase_detail_purchase FOREIGN KEY(purchase_id) REFERENCES purchases(id) ON DELETE CASCADE,
CONSTRAINT fk_purchase_detail_product FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE CASCADE
