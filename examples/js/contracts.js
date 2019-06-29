module.exports = {
  'RangeProofValidator': {
    'address': '0x2Bd63177f61D0Dc67534d48406c1849005dfE04B',
    'abi': [
      {
        'constant': true,
        'inputs': [
          {
            'name': 'lower',
            'type': 'uint256'
          },
          {
            'name': 'upper',
            'type': 'uint256'
          },
          {
            'name': 'commitment',
            'type': 'bytes'
          },
          {
            'name': 'proof',
            'type': 'bytes'
          }
        ],
        'name': 'validate',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xa61d7e3e'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'sqrt',
            'type': 'bytes'
          },
          {
            'name': 'N',
            'type': 'bytes'
          }
        ],
        'name': 'validateFloorSqrt',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x0ff4a316'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'b',
            'type': 'bytes'
          },
          {
            'name': 'N',
            'type': 'bytes'
          },
          {
            'name': 'g',
            'type': 'bytes'
          },
          {
            'name': 'h',
            'type': 'bytes'
          },
          {
            'name': 'Einv',
            'type': 'bytes'
          },
          {
            'name': 'C',
            'type': 'bytes'
          },
          {
            'name': 'D1',
            'type': 'bytes'
          },
          {
            'name': 'D2',
            'type': 'bytes'
          }
        ],
        'name': 'validateCFT',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xd1ca6c12'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'N',
            'type': 'bytes'
          },
          {
            'name': 'g',
            'type': 'bytes'
          },
          {
            'name': 'h_orInv1',
            'type': 'bytes'
          },
          {
            'name': 'h_orInv2',
            'type': 'bytes'
          },
          {
            'name': 'F',
            'type': 'bytes'
          },
          {
            'name': 'c',
            'type': 'bytes'
          },
          {
            'name': 'D',
            'type': 'bytes'
          },
          {
            'name': 'D1',
            'type': 'bytes'
          },
          {
            'name': 'D2',
            'type': 'bytes'
          },
          {
            'name': 'Einv',
            'type': 'bytes'
          },
          {
            'name': 'Finv',
            'type': 'bytes'
          }
        ],
        'name': 'validateSQ',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x976021da'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'N',
            'type': 'bytes'
          },
          {
            'name': 'g1',
            'type': 'bytes'
          },
          {
            'name': 'g2',
            'type': 'bytes'
          },
          {
            'name': 'h1',
            'type': 'bytes'
          },
          {
            'name': 'h2',
            'type': 'bytes'
          },
          {
            'name': 'Einv',
            'type': 'bytes'
          },
          {
            'name': 'Finv',
            'type': 'bytes'
          },
          {
            'name': 'c',
            'type': 'bytes'
          },
          {
            'name': 'D',
            'type': 'bytes'
          },
          {
            'name': 'D1',
            'type': 'bytes'
          },
          {
            'name': 'D2',
            'type': 'bytes'
          }
        ],
        'name': 'validateEC',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x2de10da0'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'a_inv',
            'type': 'bytes'
          },
          {
            'name': 'N',
            'type': 'bytes'
          }
        ],
        'name': 'validateModInv',
        'outputs': [
          {
            'name': '',
            'type': 'bool'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x17c91190'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'N',
            'type': 'bytes'
          },
          {
            'name': 'g',
            'type': 'bytes'
          },
          {
            'name': 'h',
            'type': 'bytes'
          },
          {
            'name': 'D1',
            'type': 'bytes'
          },
          {
            'name': 'D2',
            'type': 'bytes'
          },
          {
            'name': 'Einv',
            'type': 'bytes'
          },
          {
            'name': 'c',
            'type': 'bytes'
          }
        ],
        'name': 'restoreCommitment',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xf68f4e74'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'f',
            'type': 'bytes32'
          }
        ],
        'name': 'loadHash',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x13c3dc50'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'uint256'
          }
        ],
        'name': 'bitLength',
        'outputs': [
          {
            'name': '',
            'type': 'uint256'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xe41bbcf8'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'b',
            'type': 'bytes'
          }
        ],
        'name': 'compare',
        'outputs': [
          {
            'name': 'cmp',
            'type': 'int256'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x81e79ac6'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'uint256'
          }
        ],
        'name': 'toBigInt',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xa6203b41'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'bytes'
          }
        ],
        'name': 'bignot',
        'outputs': [
          {
            'name': '',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x98802e8f'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'b',
            'type': 'bytes'
          }
        ],
        'name': 'bigadd',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x6d8e8182'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'b',
            'type': 'bytes'
          }
        ],
        'name': 'bigsub',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xa59bc5e8'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': '_a',
            'type': 'bytes'
          },
          {
            'name': '_b',
            'type': 'bytes'
          },
          {
            'name': 'negative_b',
            'type': 'bool'
          }
        ],
        'name': 'addOrSub',
        'outputs': [
          {
            'name': 'result',
            'type': 'bytes'
          },
          {
            'name': 'cmp',
            'type': 'int256'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x42d03da9'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'bytes'
          }
        ],
        'name': 'square',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x68b54a3a'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'b',
            'type': 'bytes'
          }
        ],
        'name': 'multiply',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xa4d662d8'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'a',
            'type': 'bytes'
          },
          {
            'name': 'b',
            'type': 'bytes'
          },
          {
            'name': 'N',
            'type': 'bytes'
          }
        ],
        'name': 'modmul',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x905af814'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'bytes'
          }
        ],
        'name': 'trim',
        'outputs': [
          {
            'name': 'y',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xbcc0f6fa'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'bytes'
          },
          {
            'name': 'bitShift',
            'type': 'uint256'
          }
        ],
        'name': 'shiftBitsRight',
        'outputs': [
          {
            'name': 'y',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0xff27c476'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': 'x',
            'type': 'bytes'
          },
          {
            'name': 'n',
            'type': 'int256'
          }
        ],
        'name': 'shiftLeft',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x830058ab'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': '_x',
            'type': 'bytes'
          },
          {
            'name': '_mod',
            'type': 'bytes'
          }
        ],
        'name': 'bmod',
        'outputs': [
          {
            'name': 'ret',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x44f94bda'
      },
      {
        'constant': true,
        'inputs': [
          {
            'name': '_base',
            'type': 'bytes'
          },
          {
            'name': '_exp',
            'type': 'bytes'
          },
          {
            'name': '_mod',
            'type': 'bytes'
          }
        ],
        'name': 'modexp',
        'outputs': [
          {
            'name': '',
            'type': 'bytes'
          }
        ],
        'payable': false,
        'stateMutability': 'view',
        'type': 'function',
        'signature': '0x8b4c9e83'
      }
    ]
  }
}
