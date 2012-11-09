// THIS FILE IS AUTO-GENERATED FROM camli.js
// DO NOT EDIT.
package ui

import "time"

func init() {
	Files.Add("camli.js", "/*\n"+
		"Copyright 2011 Google Inc.\n"+
		"\n"+
		"Licensed under the Apache License, Version 2.0 (the \"License\");\n"+
		"you may not use this file except in compliance with the License.\n"+
		"You may obtain a copy of the License at\n"+
		"\n"+
		"     http://www.apache.org/licenses/LICENSE-2.0\n"+
		"\n"+
		"Unless required by applicable law or agreed to in writing, software\n"+
		"distributed under the License is distributed on an \"AS IS\" BASIS,\n"+
		"WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n"+
		"See the License for the specific language governing permissions and\n"+
		"limitations under the License.\n"+
		"*/\n"+
		"\n"+
		"// Camli namespace.\n"+
		"var Camli = {};\n"+
		"\n"+
		"var disco = null;  // TODO: kill this in favor of Camli.config.\n"+
		"\n"+
		"// innerText is not W3C compliant and does not work with firefox.\n"+
		"// textContent does not work with IE.\n"+
		"// setTextContent should work with all browsers.\n"+
		"function setTextContent(ele, text) {\n"+
		"    if (\"textContent\" in ele) {\n"+
		"        ele.textContent = text;\n"+
		"        return;\n"+
		"    }\n"+
		"    if (\"innerText\" in ele) {\n"+
		"        ele.innerText = text;\n"+
		"        return;\n"+
		"    }\n"+
		"    while (element.firstChild!==null)\n"+
		"        element.removeChild(element.firstChild);\n"+
		"    element.appendChild(document.createTextNode(text));\n"+
		"}\n"+
		"\n"+
		"// Method 1 to get discovery information (JSONP style):\n"+
		"function onConfiguration(config) {\n"+
		"    Camli.config = disco = config;\n"+
		"    console.log(\"Got config: \" + JSON.stringify(config));\n"+
		"}\n"+
		"\n"+
		"function saneOpts(opts) {\n"+
		"    if (!opts) {\n"+
		"        opts = {}\n"+
		"    }\n"+
		"    if (!opts.success) {\n"+
		"        opts.success = function() {};\n"+
		"    }\n"+
		"    if (!opts.fail) {\n"+
		"        opts.fail = function() {};\n"+
		"    }\n"+
		"    return opts;\n"+
		"}\n"+
		"\n"+
		"// Format |dateVal| as specified by RFC 3339.\n"+
		"function dateToRfc3339String(dateVal) {\n"+
		"    // Return a string containing |num| zero-padded to |length| digits.\n"+
		"    var pad = function(num, length) {\n"+
		"        var numStr = \"\" + num;\n"+
		"        while (numStr.length < length) {\n"+
		"            numStr = \"0\" + numStr;\n"+
		"        }\n"+
		"        return numStr;\n"+
		"    }\n"+
		"    return dateVal.getUTCFullYear() + \"-\" + pad(dateVal.getUTCMonth() + 1, 2) + \""+
		"-\" + pad(dateVal.getUTCDate(), 2) + \"T\" +\n"+
		"           pad(dateVal.getUTCHours(), 2) + \":\" + pad(dateVal.getUTCMinutes(), 2) "+
		"+ \":\" + pad(dateVal.getUTCSeconds(), 2) + \"Z\";\n"+
		"}\n"+
		"\n"+
		"var cachedCamliSigDiscovery;\n"+
		"\n"+
		"// opts.success called with discovery object\n"+
		"// opts.fail called with error text\n"+
		"function camliSigDiscovery(opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    if (cachedCamliSigDiscovery) {\n"+
		"        opts.success(cachedCamliSigDiscovery);\n"+
		"        return;\n"+
		"    }\n"+
		"    var cb = {};\n"+
		"    cb.success = function(sd) {\n"+
		"      cachedCamliSigDiscovery = sd;\n"+
		"      opts.success(sd);\n"+
		"    };\n"+
		"    cb.fail = opts.fail;\n"+
		"    var xhr = camliJsonXhr(\"camliSigDiscovery\", cb);\n"+
		"    xhr.open(\"GET\", Camli.config.jsonSignRoot + \"/camli/sig/discovery\", true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliDescribeBlob(blobref, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliDescribeBlob\", opts);\n"+
		"    var path = Camli.config.searchRoot + \"camli/search/describe?blobref=\" +\n"+
		"        blobref;\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function makeURL(base, map) {\n"+
		"    for (var key in map) {\n"+
		"        if (base.indexOf(\"?\") == -1) {\n"+
		"            base += \"?\";\n"+
		"        } else {\n"+
		"            base += \"&\";\n"+
		"        }\n"+
		"        base += key + \"=\" + encodeURIComponent(map[key]);\n"+
		"    }\n"+
		"    return base;\n"+
		"}\n"+
		"\n"+
		"function camliPermanodeOfSignerAttrValue(signer, attr, value, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliPermanodeOfSignerAttrValue\", opts);\n"+
		"    var path = makeURL(Camli.config.searchRoot + \"camli/search/signerattrvalue\",\n"+
		"                       { signer: signer, attr: attr, value: value });\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"// Where is the target accessed via? (paths it's at)\n"+
		"function camliPathsOfSignerTarget(signer, target, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliPathsOfSignerTarget\", opts);\n"+
		"    var path = makeURL(Camli.config.searchRoot + \"camli/search/signerpaths\",\n"+
		"                           { signer: signer, target: target });\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliGetPermanodeClaims(permanode, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliGetPermanodeClaims\", opts);\n"+
		"    var path = Camli.config.searchRoot + \"camli/search/claims?permanode=\" +\n"+
		"        permanode;\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliGetBlobContents(blobref, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    var xhr = new XMLHttpRequest();\n"+
		"    xhr.onreadystatechange = function() {\n"+
		"        if (xhr.readyState != 4) { return; }\n"+
		"        if (xhr.status != 200) {\n"+
		"            opts.fail(\"camliGetBlobContents HTTP status \" + xhr.status);\n"+
		"            return;\n"+
		"        }\n"+
		"        opts.success(xhr.responseText);\n"+
		"    };\n"+
		"    xhr.open(\"GET\", camliBlobURL(blobref), true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliBlobURL(blobref) {\n"+
		"    return Camli.config.blobRoot + \"camli/\" + blobref;\n"+
		"}\n"+
		"\n"+
		"function camliDescribeBlogURL(blobref) {\n"+
		"    return Camli.config.searchRoot + 'camli/search/describe?blobref=' + blobref;\n"+
		"}\n"+
		"\n"+
		"function camliSign(clearObj, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"\n"+
		"    camliSigDiscovery(\n"+
		"        {\n"+
		"            success: function(sigConf) {\n"+
		"                if (!sigConf.publicKeyBlobRef) {\n"+
		"                    opts.fail(\"Missing sigConf.publicKeyBlobRef\");\n"+
		"                    return;\n"+
		"                }\n"+
		"                clearObj.camliSigner = sigConf.publicKeyBlobRef;\n"+
		"                clearText = JSON.stringify(clearObj, null, 2);\n"+
		"\n"+
		"                var xhr = new XMLHttpRequest();\n"+
		"                xhr.onreadystatechange = function() {\n"+
		"                    if (xhr.readyState != 4) { return; }\n"+
		"                    if (xhr.status != 200) {\n"+
		"                        opts.fail(\"got status \" + xhr.status);\n"+
		"                        return;\n"+
		"                    }\n"+
		"                    opts.success(xhr.responseText);\n"+
		"                };\n"+
		"                xhr.open(\"POST\", sigConf.signHandler, true);\n"+
		"                xhr.setRequestHeader(\"Content-Type\", \"application/x-www-form-urle"+
		"ncoded\");\n"+
		"                xhr.send(\"json=\" + encodeURIComponent(clearText));\n"+
		"            },\n"+
		"            fail: function(errMsg) {\n"+
		"                opts.fail(errMsg);\n"+
		"            }\n"+
		"        });\n"+
		"}\n"+
		"\n"+
		"// file: File object\n"+
		"// contentsBlobRef: blob ref of file as sha1'd locally\n"+
		"// opts: fail(strMsg) success(strFileBlobRef) of the validated (or uploaded + cre"+
		"ated) file schema blob.\n"+
		"//       associating with a permanode is caller's job.\n"+
		"function camliUploadFileHelper(file, contentsBlobRef, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    if (!Camli.config.uploadHelper) {\n"+
		"        opts.fail(\"no uploadHelper available\");\n"+
		"        return\n"+
		"    }\n"+
		"\n"+
		"    var doUpload = function() {\n"+
		"        var fd = new FormData();\n"+
		"        fd.append(fd, file);\n"+
		"        var uploadCb = { fail: opts.fail };\n"+
		"        uploadCb.success = function(res) {\n"+
		"            if (res.got && res.got.length == 1 && res.got[0].fileref) {\n"+
		"                var fileblob = res.got[0].fileref;\n"+
		"                console.log(\"uploaded \" + contentsBlobRef + \" => file blob \" + fi"+
		"leblob);\n"+
		"                opts.success(fileblob);\n"+
		"            } else {\n"+
		"                opts.fail(\"failed to upload \" + file.name + \": \" + contentsBlobRe"+
		"f + \": \" + JSON.stringify(res, null, 2))\n"+
		"            }\n"+
		"        };\n"+
		"        var xhr = camliJsonXhr(\"camliUploadFileHelper\", uploadCb);\n"+
		"        xhr.open(\"POST\", Camli.config.uploadHelper);\n"+
		"        xhr.send(fd);\n"+
		"    };\n"+
		"\n"+
		"    var dupcheckCb = { fail: opts.fail };\n"+
		"    dupcheckCb.success = function(res) {\n"+
		"        var remain = res.files;\n"+
		"        var checkNext;\n"+
		"        checkNext = function() {\n"+
		"            if (remain.length == 0) {\n"+
		"                doUpload();\n"+
		"                return;\n"+
		"            }\n"+
		"            // TODO: verify filename and other file metadata in the\n"+
		"            // file json schema match too, not just the contents\n"+
		"            var checkFile = remain.shift();\n"+
		"            console.log(\"integrity checking the reported dup \" + checkFile);\n"+
		"\n"+
		"            var vcb = {};\n"+
		"            vcb.fail = function(xhr) {\n"+
		"                console.log(\"integrity checked failed on \" + checkFile);\n"+
		"                checkNext();\n"+
		"            };\n"+
		"            vcb.success = function(xhr) {\n"+
		"                if (xhr.getResponseHeader(\"X-Camli-Contents\") == contentsBlobRef)"+
		" {\n"+
		"                    console.log(\"integrity checked passed on \" + checkFile + \"; u"+
		"sing it.\");\n"+
		"                    opts.success(checkFile);\n"+
		"                } else {\n"+
		"                    checkNext();\n"+
		"                }\n"+
		"            };\n"+
		"            var xhr = camliXhr(\"headVerifyFile\", vcb);\n"+
		"            xhr.open(\"HEAD\", Camli.config.downloadHelper + checkFile + \"/?verifyc"+
		"ontents=\" + contentsBlobRef, true);\n"+
		"            xhr.send();\n"+
		"        };\n"+
		"        checkNext();\n"+
		"    };\n"+
		"    camliFindExistingFileSchemas(contentsBlobRef, dupcheckCb);\n"+
		"}\n"+
		"\n"+
		"function camliUploadString(s, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    var blobref = \"sha1-\" + Crypto.SHA1(s);\n"+
		"    var parts = [s];\n"+
		"\n"+
		"    var bb = new Blob(parts);\n"+
		"\n"+
		"    var fd = new FormData();\n"+
		"    fd.append(blobref, bb);\n"+
		"\n"+
		"    var uploadCb = {};\n"+
		"    uploadCb.success = function(resj) {\n"+
		"        // TODO: check resj.received[] array.\n"+
		"        opts.success(blobref);\n"+
		"    };\n"+
		"    uploadCb.fail = opts.fail;\n"+
		"    var xhr = camliJsonXhr(\"camliUploadString\", uploadCb);\n"+
		"    // TODO: hack, hard-coding the upload URL here.\n"+
		"    // Change the spec now that App Engine permits 32 MB requests\n"+
		"    // and permit a PUT request on the sha1?  Or at least let us\n"+
		"    // specify the well-known upload URL?  In cases like this, uploading\n"+
		"    // a new permanode, it's silly to even stat.\n"+
		"    xhr.open(\"POST\", Camli.config.blobRoot + \"camli/upload\")\n"+
		"    xhr.send(fd);\n"+
		"}\n"+
		"\n"+
		"function camliCreateNewPermanode(opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"     var json = {\n"+
		"         \"camliVersion\": 1,\n"+
		"         \"camliType\": \"permanode\",\n"+
		"         \"random\": \"\"+Math.random()\n"+
		"     };\n"+
		"     camliSign(json, {\n"+
		"                   success: function(got) {\n"+
		"                       camliUploadString(\n"+
		"                           got,\n"+
		"                           {\n"+
		"                               success: opts.success,\n"+
		"                               fail: function(msg) {\n"+
		"                                   opts.fail(\"upload permanode fail: \" + msg);\n"+
		"                               }\n"+
		"                           });\n"+
		"                   },\n"+
		"                   fail: function(msg) {\n"+
		"                       opts.fail(\"sign permanode fail: \" + msg);\n"+
		"                   }\n"+
		"               });\n"+
		"}\n"+
		"\n"+
		"// Returns the first value from the query string corresponding to |key|.\n"+
		"// Returns null if the key isn't present.\n"+
		"function getQueryParam(key) {\n"+
		"    var params = document.location.search.substring(1).split('&');\n"+
		"    for (var i = 0; i < params.length; ++i) {\n"+
		"        var parts = params[i].split('=');\n"+
		"        if (parts.length == 2 && decodeURIComponent(parts[0]) == key)\n"+
		"            return decodeURIComponent(parts[1]);\n"+
		"    }\n"+
		"    return null;\n"+
		"}\n"+
		"\n"+
		"function camliGetRecentlyUpdatedPermanodes(opts) {\n"+
		"    // opts.thumbnails is the maximum size of the thumbnails we want,\n"+
		"    // or 0 if no thumbnail.\n"+
		"    var path = \"\";\n"+
		"    if (opts.thumbnails != null) {\n"+
		"        path = makeURL(Camli.config.searchRoot + \"camli/search/recent\",\n"+
		"            {thumbnails: opts.thumbnails});\n"+
		"    } else {\n"+
		"        path = makeURL(Camli.config.searchRoot + \"camli/search/recent\");\n"+
		"    }\n"+
		"    var xhr = camliJsonXhr(\"camliGetRecentlyUpdatedPermanodes\", opts);\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliGetPermanodesWithAttr(signer, attr, value, fuzzy, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliGetPermanodesWithAttr\", opts);\n"+
		"    var path = makeURL(Camli.config.searchRoot + \"camli/search/permanodeattr\",\n"+
		"                       { signer: signer, attr: attr, value: value, fuzzy: fuzzy }"+
		");\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"function camliXhr(name, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    var xhr = new XMLHttpRequest();\n"+
		"    xhr.onreadystatechange = function() {\n"+
		"        if (xhr.readyState != 4) { return; }\n"+
		"        if (xhr.status == 200) {\n"+
		"            opts.success(xhr);\n"+
		"        } else {\n"+
		"            opts.fail(name + \": expected status 200; got \" + xhr.status + \": \" + "+
		"xhr.responseText);\n"+
		"        }\n"+
		"    };\n"+
		"    return xhr;\n"+
		"}\n"+
		"\n"+
		"function camliJsonXhr(name, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    var xhr = new XMLHttpRequest();\n"+
		"    xhr.onreadystatechange = function() {\n"+
		"        if (xhr.readyState != 4) { return; }\n"+
		"        if (xhr.status != 200) {\n"+
		"            try {\n"+
		"                var resj = JSON.parse(xhr.responseText);\n"+
		"                opts.fail(name + \": expected status 200; got \" + xhr.status + \": "+
		"\" + resj.error);\n"+
		"            } catch(x) {\n"+
		"                opts.fail(name + \": expected status 200; got \" + xhr.status + \": "+
		"\" + xhr.responseText);\n"+
		"            }\n"+
		"            return;\n"+
		"        }\n"+
		"        var resj;\n"+
		"        try {\n"+
		"            resj = JSON.parse(xhr.responseText);\n"+
		"        } catch(x) {\n"+
		"            opts.fail(name + \": error parsing JSON in response: \" + xhr.responseT"+
		"ext);\n"+
		"            return\n"+
		"        }\n"+
		"        if (resj.error) {\n"+
		"            opts.fail(resj.error);\n"+
		"        } else {\n"+
		"            opts.success(resj);\n"+
		"        }\n"+
		"    };\n"+
		"    return xhr;\n"+
		"}\n"+
		"\n"+
		"function camliFindExistingFileSchemas(wholeDigestRef, opts) {\n"+
		"    var xhr = camliJsonXhr(\"camliFindExistingFileSchemas\", opts);\n"+
		"    var path = Camli.config.searchRoot + \"camli/search/files?wholedigest=\" +\n"+
		"        wholeDigestRef;\n"+
		"    xhr.open(\"GET\", path, true);\n"+
		"    xhr.send();\n"+
		"}\n"+
		"\n"+
		"// Returns true if the passed-in string might be a blobref.\n"+
		"function isPlausibleBlobRef(blobRef) {\n"+
		"    return /^\\w+-[a-f0-9]+$/.test(blobRef);\n"+
		"}\n"+
		"\n"+
		"function linkifyBlobRefs(schemaBlob) {\n"+
		"    var re = /(\\w{3,6}-[a-f0-9]{30,})/g;\n"+
		"    return schemaBlob.replace(re, \"<a href='./?b=$1'>$1</a>\");\n"+
		"}\n"+
		"\n"+
		"// Helper function for camliNewSetAttributeClaim() (and eventually, for\n"+
		"// similar functions to add or delete attributes).\n"+
		"function changeAttribute(permanode, claimType, attribute, value, opts) {\n"+
		"    opts = saneOpts(opts);\n"+
		"    var json = {\n"+
		"        \"camliVersion\": 1,\n"+
		"        \"camliType\": \"claim\",\n"+
		"        \"permaNode\": permanode,\n"+
		"        \"claimType\": claimType,\n"+
		"        \"claimDate\": dateToRfc3339String(new Date()),\n"+
		"        \"attribute\": attribute,\n"+
		"        \"value\": value\n"+
		"    };\n"+
		"    camliSign(json, {\n"+
		"        success: function(signedBlob) {\n"+
		"            camliUploadString(signedBlob, {\n"+
		"                success: opts.success,\n"+
		"                fail: function(msg) {\n"+
		"                    opts.fail(\"upload \" + claimType + \" fail: \" + msg);\n"+
		"                }\n"+
		"            });\n"+
		"        },\n"+
		"        fail: function(msg) {\n"+
		"            opts.fail(\"sign \" + claimType + \" fail: \" + msg);\n"+
		"        }\n"+
		"    });\n"+
		"}\n"+
		"\n"+
		"function camliBlobTitle(pn, des) {\n"+
		"    return _camliBlobTitleOrThumb(pn, des, 0, 0);\n"+
		"}\n"+
		"\n"+
		"function camliBlobThumbnail(pn, des, width, height) {\n"+
		"    return _camliBlobTitleOrThumb(pn, des, width, height);\n"+
		"}\n"+
		"\n"+
		"// pn: permanode to find a good title of\n"+
		"// jdes: describe response of root permanode\n"+
		"// w, h: if both of them are non-zero, returns html of an wxh size <img> thumbnai"+
		"l, not a title.\n"+
		"function _camliBlobTitleOrThumb(pn, des, w, h) {\n"+
		"    var d = des[pn];\n"+
		"    if (!d) {\n"+
		"        return pn;\n"+
		"    }\n"+
		"    if (d.camliType == \"file\" && d.file && d.file.fileName) {\n"+
		"        var fileName = d.file.fileName\n"+
		"        if (w != 0 && h != 0 && d.file.mimeType && d.file.mimeType.indexOf(\"image"+
		"/\") == 0) {\n"+
		"            var img = \"<img src='./thumbnail/\" + pn + \"/\" +\n"+
		"            fileName.replace(/['\"<>\\?&]/g, \"\") + \"?mw=\" + w + \"&mh=\" + h + \"'>\";\n"+
		"            return img;\n"+
		"        }\n"+
		"        return fileName;\n"+
		"    }\n"+
		"    if (d.permanode) {\n"+
		"        var attr = d.permanode.attr;\n"+
		"        if (!attr) {\n"+
		"            return pn;\n"+
		"        }\n"+
		"        if (attr.title) {\n"+
		"            return attr.title[0];\n"+
		"        }\n"+
		"        if (attr.camliContent) {\n"+
		"            return _camliBlobTitleOrThumb(attr.camliContent[0], des, w, h);\n"+
		"        }\n"+
		"    }\n"+
		"    return pn;\n"+
		"}\n"+
		"\n"+
		"// Create and upload a new set-attribute claim.\n"+
		"function camliNewSetAttributeClaim(permanode, attribute, value, opts) {\n"+
		"    changeAttribute(permanode, \"set-attribute\", attribute, value, opts);\n"+
		"}\n"+
		"\n"+
		"// Create and upload a new add-attribute claim.\n"+
		"function camliNewAddAttributeClaim(permanode, attribute, value, opts) {\n"+
		"    changeAttribute(permanode, \"add-attribute\", attribute, value, opts);\n"+
		"}\n"+
		"\n"+
		"// Create and upload a new del-attribute claim.\n"+
		"function camliNewDelAttributeClaim(permanode, attribute, value, opts) {\n"+
		"    changeAttribute(permanode, \"del-attribute\", attribute, value, opts);\n"+
		"}\n"+
		"\n"+
		"", time.Unix(0, 1352486005211518273))
}
