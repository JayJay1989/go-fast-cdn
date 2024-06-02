import { useState } from "react";
import UploadPreview from "./upload-preview";

const FilesInput: React.FC<{
  fileRef: React.RefObject<HTMLInputElement>;
}> = ({ fileRef }) => {
  const [fileNames, setFileNames] = useState<string[]>([]);

  const getFileNames = () => {
    if (fileRef.current?.files != null) {
      let names = [];

      for (let file of fileRef.current.files) {
        names.push(file.name);
      }

      setFileNames(names);
    }
  };

  return (
    <div className="my-8 flex max-h-[691px] flex-col justify-center items-center h-full">
      <div className="w-full h-full mb-4 flex justify-center items-center overflow-hidden">
        <UploadPreview fileNames={fileNames} type="files"/>
      </div>
      <label htmlFor="file" className="flex flex-col">
        Select File
        <input
          onChange={getFileNames}
          type="file"
          accept="application/vnd.rar, application/zip, application/x-7z-compressed, XML, application/octet-stream, application/vnd.microsoft.portable-executable, application/x-ms-installer"
          multiple
          name="file"
          id="file"
          aria-label="Select file"
          ref={fileRef}
        />
      </label>
    </div>
  );
};

export default FilesInput;
