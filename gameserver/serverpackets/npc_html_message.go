package serverpackets

import "l2gogameserver/gameserver/models"

func NpcHtmlMessage(client *models.Client) {
	client.Buffer.WriteSingleByte(0x19)

	client.Buffer.WriteD(33)
	client.Buffer.WriteS("<html><title>Shop</title><body>\n<center>\n<table width=260>\n<tr><td width=40><button value=\"Main\" action=\"bypass -h admin_admin\" width=40 height=15 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"></td>\n<td width=180><center>GM Shop - Territory Items</center></td>\n<td width=40><button value=\"Back\" action=\"bypass -h admin_html gmshops.htm\" width=40 height=15 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"></td>\n</tr></table>\n</center>\n<br><br>\n<center>\n<button action=\"bypass -h admin_buy 9988\" value=\"Territory Items\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9990\" value=\"Territory Weapons\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9987\" value=\"Territory Jewels\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9991\" value=\"Territory Wards\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9992\" value=\"Territory Flags\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9970\" value=\"Mercenary Transforms\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n<button action=\"bypass -h admin_buy 9971\" value=\"Disguise Scroll\" width=150 height=21 back=\"L2UI_CT1.Button_DF_Down\" fore=\"L2UI_CT1.Button_DF\"><br1>\n</center>\n</body></html>")
	client.Buffer.WriteD(0)
	client.SaveAndCryptDataInBufferToSend(true)
}
