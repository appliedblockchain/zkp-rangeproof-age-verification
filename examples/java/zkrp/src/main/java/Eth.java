import com.ing.blockchain.zk.dto.ClosedRange;
import com.ing.blockchain.zk.dto.Commitment;
import com.ing.blockchain.zk.dto.SecretOrderGroup;
import com.ing.blockchain.zk.util.ExportUtil;

import java.math.BigInteger;

public class Eth { // government or KYC provider
  public static byte[] exportSolidity(SecretOrderGroup group, BigInteger commitmentValue) {
    Commitment commitment = new Commitment(group, commitmentValue);
    byte[] solidityCommitment = ExportUtil.exportForEVM(commitment);

    Utils.print("Ethereum solidity output:");
    Utils.print(solidityCommitment);

    return solidityCommitment;
  }
}

