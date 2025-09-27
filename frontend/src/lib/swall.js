import Swal from "sweetalert2";

export const success = async (message) => {
  Swal.fire({
    icon: "success",
    title: `${message} Telah Berhasil`,
    // text: message,
  });
};

export const failed = async (message) => {
  Swal.fire({
    icon: "error",
    title: "Oops...",
    text: message,
  });
};

const Toast = Swal.mixin({
  toast: true,
  position: "top-end",
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: (toast) => {
    toast.onmouseenter = Swal.stopTimer;
    toast.onmouseleave = Swal.resumeTimer;
  },
});

export const succesLogin = async (message) => {
  Toast.fire({
    icon: "success",
    title: "Login in successfully",
  });
};
