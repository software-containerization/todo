import axios from 'axios';

const API_URL = 'API_BASE_URL'

async function createItem(item) {
    const res = await axios.post(`${API_URL}/item`, {
        title: item.title
    });

    return res.data;
}

async function deleteItem(id) {
    const message = await axios.delete(`${API_URL}/item/${id}`);
    return message;
}

async function updateItem(id, payload) {
    const { data: newItem } = await axios.put(`${API_URL}/item/${id}`, payload);
    return newItem;
}

async function getAllItems() {
    const { data: items } = await axios.get(`${API_URL}/items`);
    return items;
}

export default { createItem, deleteItem, updateItem, getAllItems };
