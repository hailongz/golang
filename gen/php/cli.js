"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const less = __importStar(require("../Less"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
let url = require("url");
function getType(fd, ns) {
    let v = "String";
    switch (fd.type) {
        case less.FieldType.INT32:
            v = "int";
            break;
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            v = "float";
            break;
        case less.FieldType.INT64:
            v = "int";
            break;
        case less.FieldType.BOOLEAN:
            v = "boolean";
            break;
        case less.FieldType.OBJECT:
        case less.FieldType.ENUM:
            if (fd.typeSymbol === undefined) {
                v = "any";
            }
            else {
                v = ns ? '\\' + ns + '\\' + fd.typeSymbol : fd.typeSymbol;
            }
            break;
    }
    if (fd.isArray) {
        v = v + '[]';
    }
    return v;
}
function mkdirs(dir) {
    if (fs.existsSync(dir)) {
        return;
    }
    mkdirs(path.dirname(dir));
    fs.mkdirSync(dir);
}
function walk(basePath, outfile, ns) {
    var objectSet = {};
    var code = [];
    var packageSet = {};
    code.push("<?php\n");
    code.push("namespace " + ns + ";\n\n\n");
    less.walk(basePath, (v) => {
        for (let object of v.enums) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                for (let fd of object.items) {
                    if (fd.title) {
                        code.push("\t/**\n");
                        code.push("\t * ");
                        if (object.title) {
                            code.push(object.title);
                            code.push(" - ");
                        }
                        code.push(fd.title);
                        code.push("\n\t */\n");
                    }
                    else if (object.title) {
                        code.push("\t/**\n");
                        code.push("\t * ");
                        code.push(object.title);
                        code.push("\n\t */\n");
                    }
                    code.push(`define(${object.name.toLocaleUpperCase()}_${fd.name},${JSON.stringify(fd.value)})\n`);
                }
                code.push("\n");
            }
        }
        for (let object of v.objects) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                if (object.title) {
                    code.push("/**\n");
                    code.push(" * ");
                    code.push(object.title);
                    code.push("\n */\n");
                }
                code.push("class ");
                code.push(object.name);
                code.push(" {\n\n");
                for (let fd of object.fields) {
                    code.push("\t/**\n");
                    code.push("\t * ");
                    code.push((fd.title || '').replace(/\n/g, ' '));
                    code.push("\n\t * @var ");
                    code.push(getType(fd, ns));
                    if (fd.required) {
                        code.push("\n\t * @required ");
                    }
                    code.push("\n\t */\n");
                    code.push("\tpublic $");
                    code.push(fd.name);
                    code.push(";\n\n");
                }
                code.push("}\n\n");
            }
        }
        {
            let dir = path.dirname(v.name);
            let vs = packageSet[dir];
            if (vs === undefined) {
                packageSet[dir] = [v];
            }
            else {
                vs.push(v);
            }
        }
    });
    for (let p in packageSet) {
        let vs = packageSet[p];
        let basename = p.replace(/\//g, '_');
        if (p == ".") {
            basename = '';
        }
        else if (basename == '_') {
            basename = '';
        }
        else {
            basename = basename + '_';
        }
        for (let v of vs) {
            lessCode(code, v, basename, ns);
        }
    }
    fs.writeFileSync(outfile, code.join(''));
}
exports.walk = walk;
function lessCode(vs, v, basename, ns) {
    {
        let data = "any";
        for (let fd of v.response.fields) {
            if (fd.name == "data" && fd.typeSymbol !== undefined) {
                data = ns ? '\\' + ns + '\\' + fd.typeSymbol : fd.typeSymbol;
                break;
            }
        }
        let n = path.basename(v.name);
        if (n == 'in') {
            n = "In";
        }
        if (n == 'for') {
            n = "For";
        }
        if (n == 'If') {
            n = "If";
        }
        if (n == 'delete') {
            n = "Delete";
        }
        vs.push(`class ${basename}${n}_Task {\n`);
        for (let fd of v.request.fields) {
            vs.push("\t/**\n");
            vs.push("\t * ");
            vs.push((fd.title || '').replace(/\n/g, ' '));
            vs.push("\n\t * @var ");
            vs.push(getType(fd, ns));
            if (fd.required) {
                vs.push("\n\t * @required");
            }
            vs.push("\n\t */\n");
            vs.push("\tpublic $");
            vs.push(fd.name);
            vs.push(";\n\n");
        }
        vs.push('}\n\n');
        vs.push("/**\n");
        vs.push(" * ");
        vs.push(v.request.title || '');
        vs.push("\n * @param ");
        vs.push(`String $name`);
        vs.push("\n * @param ");
        vs.push(`${ns ? '\\' + ns + '\\' : ''}${basename}${n}_Task $task`);
        vs.push("\n * @return ");
        vs.push(data);
        vs.push("\n */\n");
        vs.push(`function ${basename}${n}($name,$task) {\n`);
        vs.push(`\treturn \\kk\\send($name,${JSON.stringify(v.request.method || 'POST')},${JSON.stringify(v.name + '.json')},$task);\n`);
        vs.push('}\n\n');
    }
}
