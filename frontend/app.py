from flask import Flask, render_template, request, jsonify
import requests

app = Flask(__name__)

# Base URL for the backend API
BASE_URL = "http://localhost:80"  # 替换为实际后端服务地址


@app.route("/")
def index():
    """Home page with links to various forms."""
    return render_template("index.html")


# ** Get Metadata **
@app.route("/get_metadata", methods=["GET", "POST"])
def get_metadata():
    if request.method == "POST":
        # Get data from form
        sn = request.form.get("sn")
        # Call the backend API
        response = requests.get(f"{BASE_URL}/metadata", params={"sn": sn})
        return jsonify(response.json())
    # Render the form page
    return render_template("form_pages/get_metadata.html")


# ** Update Metadata **
@app.route("/update_metadata", methods=["GET", "POST"])
def update_metadata():
    if request.method == "POST":
        # Get data from form
        sn = request.form.get("sn")
        params = request.form.get("params")
        # Call the backend API
        response = requests.post(
            f"{BASE_URL}/metadata",
            json={"sn": sn, "params": eval(params)},  # Convert string to dict
        )
        return jsonify(response.json())

    # Render the form page
    return render_template("form_pages/update_metadata.html")


# ** Get Metrics **
@app.route("/get_metrics", methods=["GET", "POST"])
def get_metrics():
    if request.method == "POST":
        # Get data from form
        sn = request.form.get("sn")
        start_time = request.form.get("start_time")
        end_time = request.form.get("end_time")
        # Call the backend API
        response = requests.get(
            f"{BASE_URL}/metrics",
            params={"sn": sn, "start_time": start_time, "end_time": end_time},
        )
        return jsonify(response.json())

    # Render the form page
    return render_template("form_pages/get_metrics.html")


# ** Other Routes (for other APIs) **
# You can add similar routes for other APIs like GenerateDemoMetadata, RandomUpdateMetadata, etc.


if __name__ == "__main__":
    app.run(debug=True)