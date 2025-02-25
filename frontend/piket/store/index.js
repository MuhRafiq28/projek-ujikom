export const state = () => ({
  user: {
    role: "user", // Default jika tidak ada data
  },
});

export const mutations = {
  SET_ROLE(state, role) {
    console.log("Set role di Vuex:", role); // Debugging
    state.user.role = role;
  },
};

export const actions = {
  fetchUserRole({ commit }) {
    const role = localStorage.getItem("userRole") || "user"; // Ambil dari localStorage
    console.log("Role dari localStorage:", role); // Debugging
    commit("SET_ROLE", role);
  },
};
export const strict = false;

