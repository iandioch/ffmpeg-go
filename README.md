This project is a small Go script which uses `ffmpeg` to create a GIF of a section of a video hosted at a given URL.

# Setup

1. Requires working commands `ffmpeg` and `uuidgen`.
2. Run `mkdir out` in this repo to create the directory for the output GIFs.

# Usage

If the server is running at `http://server.com:1234`, then call `http://server.com:1234/extract_yt?start=0&dur=10&video=VIDEO_URL`.

Ensure that the `VIDEO_URL` argument is properly URL-encoded.

You will then be redirected to a URL like `http://server.com/gif/883ea10e-d9da-4f56-8ca1-09d8a8ac6ff8.gif`.

# Maintenance

Output GIFs will be put in the `out/` directory created in step 2 of Setup above. By default, these are not deleted, so the directory may need to be cleaned regularly.
