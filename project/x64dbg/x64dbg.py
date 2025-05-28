# /// script
# requires-python = ">=3.10"
# dependencies = [
#     "requests>=2,<3",
#     "mcp>=1.2.0,<2",
# ]
# ///

import sys
import requests

from mcp.server.fastmcp import FastMCP

DEFAULT_X64DBG_SERVER = "http://127.0.0.1:8888/"
x64dbg_server_url = sys.argv[1] if len(sys.argv) > 1 else DEFAULT_X64DBG_SERVER

mcp = FastMCP("x64dbg-mcp")

def safe_get(endpoint: str, params: dict = None) -> list:
    """
    Perform a GET request with optional query parameters.
    """
    if params is None:
        params = {}

    url = f"{x64dbg_server_url}/{endpoint}"

    try:
        response = requests.get(url, params=params, timeout=15)
        response.encoding = 'utf-8'
        if response.ok:
            return response.text.splitlines()
        else:
            return [f"Error {response.status_code}: {response.text.strip()}"]
    except Exception as e:
        return [f"Request failed: {str(e)}"]

def safe_post(endpoint: str, data: dict | str) -> str:
    try:
        if isinstance(data, dict):
            response = requests.post(f"{x64dbg_server_url}/{endpoint}", data=data, timeout=5)
        else:
            response = requests.post(f"{x64dbg_server_url}/{endpoint}", data=data.encode("utf-8"), timeout=5)
        response.encoding = 'utf-8'
        if response.ok:
            return response.text.strip()
        else:
            return f"Error {response.status_code}: {response.text.strip()}"
    except Exception as e:
        return f"Request failed: {str(e)}"

@mcp.tool()
def Is_Debugging(Dbgcheck: str) -> str:
    """
    Find if currently debugging
    """
    return safe_get("IsDebugActive", {"dbgcheck": Dbgcheck})

@mcp.tool()
def ExeConsoleCmd(Command: str) -> list:
    """
    Execute a console command
    """
    return safe_get("ExeConsoleCmd", {"Command": Command})

@mcp.tool()
def MemoryBase(addr: str) -> dict:
    """
    Find the base address and size of a module containing the given address
    
    Parameters:
        addr: Memory address (in hex format, e.g. "0x7FF12345")
    
    Returns:
        Dictionary containing base_address and size of the module
    """
    try:
        # Print detailed debugging info
        print(f"Calling MemoryBase with address: {addr}")
        
        # Make the request
        response = safe_get("/FindMemBase", {"addr": addr})
        print(f"Raw response: '{response}'")
        
        # Return the raw response if we can't parse it properly
        if not response:
            print("No response received")
            return {"error": "No response received"}
            
        # Try to parse the response
        if "," in response:
            parts = response.split(",")
            print(f"Split parts: {parts}")
            return {
                "base_address": parts[0],
                "size": int(parts[1])
            }
        else:
            # If we can't parse it as expected, just return the raw response
            print("Response format not as expected")
            return {"raw_response": response}
            
    except Exception as e:
        # Return the exception info
        print(f"Exception in MemoryBase: {str(e)}")
        return {"error": str(e)}
@mcp.tool()
def GetModuleList() -> list:
    """
    Get a list of all loaded modules in the debugger
    
    Returns:
        List of dictionaries containing module information (name, base, size, entry, sectionCount, path)
    """
    response = safe_get("/GetModuleList")
    
    # Parse the JSON response
    if response:
        try:
            # If response is already parsed as a list
            if isinstance(response, list):
                return response
            
            # Otherwise, try to parse it as JSON
            import json
            return json.loads(response)
        except Exception as e:
            return {"error": f"Failed to parse response: {str(e)}", "raw_response": response}
    
    return None
    
if __name__ == "__main__":
    mcp.run()