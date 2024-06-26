import axios from "axios";
import { SetStateAction } from "jotai";
import toast from "react-hot-toast";
import File from "../types/file";
import { SetAtom } from "@/types/setAtom";

export const getFiles = (
  type: "images" | "documents" | "files",
  setFiles: SetAtom<[SetStateAction<File[]>], void>,
) => {
  toast.loading("Loading files...");
  axios
    .get(`/api/cdn/${type === "images" ? "image" : type === "documents" ? "doc" : "file"}/all`)
    .then((res) => {
      res.data != null && setFiles(res.data)
      toast.dismiss();
    })
    .catch((err: Error) => {
      toast.dismiss();
      toast.error(err.message);
      console.log(err);
    })
}

